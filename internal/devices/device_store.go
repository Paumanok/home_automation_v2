package devices

import (
	"context"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type deviceStore struct {
	pdqdriver *pgxpool.Pool
	tableName string
}



func (ds *deviceStore) Create(ctx context.Context, d *Device) error {


	return nil
}


func NewStore(pdqdriver *pgxpool.Pool) (*deviceStore, error) {
	return &deviceStore{
		pdqdriver: pdqdriver,
		tableName: "Devices",
	}, nil
}