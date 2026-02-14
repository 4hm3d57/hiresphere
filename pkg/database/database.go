package database



import (
	"context"
	"time"
	"hire/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)


func New(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
