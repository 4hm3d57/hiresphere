package database



import (
	"context"
	"hire/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)


func New(cfg *config.Config) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), cfg.DB_URL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
