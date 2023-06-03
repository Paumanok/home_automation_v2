package measurements

import (
	"context"
	"fmt"
	"strings"
	//"github.com/jackc/pgx/v5"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type measurementStore struct {
	pdqdriver *pgxpool.Pool
	sq squirrel.StatementBuilderType
	tableName string
}

func (ms *measurementStore) Create(ctx context.Context, m *Measurement) error {
	query, args, err := ms.sq.Insert(ms.tableName).SetMap(map[string]interface{}{
			"mac": m.MAC,
			"temp": m.Temp,
			"humidity": m.Humidity,
			"pressure": m.Pressure,
			"pm25": m.PM25,
			"createdat": m.CreatedAt,
	}).ToSql()
	if err != nil {
		fmt.Println(err)
	}
	_, err = ms.pdqdriver.Exec(ctx, query, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			fmt.Println("measurement_store: duplicate error")
			return err
		}
		fmt.Println("measurement_store: internal error")
		fmt.Println(err)
		return err
	}
	return nil
}

func (ms *measurementStore) GetAllMeasurements(ctx context.Context) ([]Measurement, error){
	query, args, err := ms.sq.Select(
		"mac",
		"temp",
		"humidity",
		"pressure",
		"pm25",
		"createdat",
		).From(ms.tableName).ToSql()

	if err != nil {
		return nil, err
	}
	rows, err := ms.pdqdriver.Query(ctx, query, args...)
	var meas []Measurement

	for rows.Next() {
		var cur_meas Measurement
		err := rows.Scan(
			&cur_meas.MAC,
			&cur_meas.Temp,
			&cur_meas.Humidity,
			&cur_meas.Pressure,
			&cur_meas.PM25,
			&cur_meas.CreatedAt,
		)
		if err != nil {
			fmt.Println("scan error")	
			return nil, err
		}
		fmt.Printf("%+v\n", cur_meas)

		meas = append(meas, cur_meas)
	}

	return meas,nil
}

func (ms *measurementStore) GetByMAC(ctx context.Context, mac string) ([]Measurement, error){
	query, args, err := ms.sq.Select(
		"mac",
		"temp",
		"humidity",
		"pressure",
		"pm25",
		"createdAt",
	).From(
		ms.tableName,
	).Where(
		squirrel.Eq{
			"mac":mac,
		},
	).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := ms.pdqdriver.Query(ctx, query, args...)

	var meas []Measurement

	for rows.Next() {
		var cur_meas Measurement
		err :=rows.Scan(
			&cur_meas.MAC,
			&cur_meas.Temp,
			&cur_meas.Humidity,
			&cur_meas.Pressure,
			&cur_meas.PM25,
			&cur_meas.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		meas = append(meas, cur_meas)

	}

	return meas,nil
}

func NewStore(pdqdriver *pgxpool.Pool) (*measurementStore, error) {
	return &measurementStore{
		pdqdriver: pdqdriver,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		tableName: "Measurements",
	}, nil
}
