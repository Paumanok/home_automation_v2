package measurements

import (
	//"context"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type measurementStore struct {
	pdqdriver *pgxpool.Pool
	tableName string
}

func (ms *measurementStore) Create(ctx context.Context, m *Measurement) error {
	return nil
}

func NewStore(pdqdriver *pgxpool.Pool) (*measurementStore, error) {
	return &measurementStore{
		pdqdriver: pdqdriver,
		tableName: "Measurements",
	}, nil
}