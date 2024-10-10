package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/Odvin/go-live-reloading/config"
	_ "github.com/lib/pq"
)

type DbAdapter struct {
	DB *sql.DB
}

func Init(cfg *config.Config) (*DbAdapter, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)

	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &DbAdapter{DB: db}, nil
}
