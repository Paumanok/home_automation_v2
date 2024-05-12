package devices

import (
	"context"
	"fmt"
	"strings"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Masterminds/squirrel"
)


type deviceStore struct {
	pdqdriver *pgxpool.Pool
	sq squirrel.StatementBuilderType
	tableName string
}


func (ds *deviceStore) Create(ctx context.Context, d *Device) error {
	query, args, err := ds.sq.Insert(ds.tableName).SetMap(map[string]interface{}{
		"nickname": d.Nickname,
		"mac": d.MAC, 
		"humiditycomp": d.HumidityComp,
		"temperaturecomp": d.TemperatureComp,
	}).ToSql()
	if err != nil {
		fmt.Println(err)
	}
	_, err = ds.pdqdriver.Exec(ctx, query, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			fmt.Println("device_store: duplicate error")
			return err
		}
		fmt.Println("device_store: internal error")
		fmt.Println(err)
		return err
	}
	return nil
}

func (ds *deviceStore) Update(ctx context.Context, d *Device, mac string) error {
	query, args, err := ds.sq.Update(ds.tableName).SetMap(
		map[string]interface{}{
			"nickname": d.Nickname,
			"mac": d.MAC, 
			"humiditycomp": d.HumidityComp,
			"temperaturecomp": d.TemperatureComp,
	}).Where(
		squirrel.Eq{
			"mac" : mac, //i should worry about cap vs lowercase somewhere
			},
		).ToSql()
	if err != nil {
		fmt.Println(err)
	}

	_, err = ds.pdqdriver.Exec(ctx, query, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			fmt.Println("device_store: duplicate error")
			return err
		}
		fmt.Println("device_store: internal error")
		fmt.Println(err)
		return err
	}
	return nil

}

func (ds *deviceStore) GetDeviceByMac(ctx context.Context, mac string) (*Device, error){
	query, args, err := ds.sq.Select(
		"nickname",
		"mac",
		"humiditycomp",
		"temperaturecomp",
	).From(
		ds.tableName,
		).Where(
			squirrel.Eq{
				"mac" : mac, //i should worry about cap vs lowercase somewhere
			},
		).ToSql()

		if err != nil {
			return nil, err
		}

		row := ds.pdqdriver.QueryRow(ctx, query, args...)
		
		var dev Device
		
		err = row.Scan(
			&dev.Nickname,
			&dev.MAC, 
			&dev.HumidityComp,
			&dev.TemperatureComp,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &dev, nil

}

func (ds *deviceStore) GetDevices(ctx context.Context) ([]Device, error) {
	query, args, err := ds.sq.Select(
		"nickname",
		"mac",
		"humiditycomp",
		"temperaturecomp",
		).From(ds.tableName).ToSql()

	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	rows, err := ds.pdqdriver.Query(ctx, query, args...)
	var devs []Device
	for rows.Next() {
		var cur_dev Device
		err := rows.Scan(
			&cur_dev.Nickname,
			&cur_dev.MAC,
			&cur_dev.HumidityComp,
			&cur_dev.TemperatureComp,
		)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("%+v", cur_dev)
		devs = append(devs, cur_dev)
	}
	return devs, nil
}

func NewStore(pdqdriver *pgxpool.Pool) (*deviceStore, error) {
	return &deviceStore{
		pdqdriver: pdqdriver,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		tableName: "devices",
	}, nil
}
