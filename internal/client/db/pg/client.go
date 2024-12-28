package pg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solumD/WBTech_L0/internal/client/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

// New returns new Postgres client
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
