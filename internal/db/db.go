package db

import (
	"github.com/aliyasirnac/goBackendBoilerplate/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Config) (*gorm.DB, error) {
	p := config.NewPostgres(cfg.Database)
	db, err := gorm.Open(postgres.Open(p.Dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
