package app

import (
	"context"

	"github.com/aliyasirnac/goBackendBoilerplate/internal/app/api"
	"github.com/aliyasirnac/goBackendBoilerplate/internal/config"
	"github.com/aliyasirnac/goBackendBoilerplate/internal/db"
	"github.com/aliyasirnac/goBackendBoilerplate/internal/loggerx"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg    *config.Config
	logger logrus.FieldLogger
}

func New(cfg *config.Config) *App {
	logger := loggerx.New(cfg.App.Log)
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.logger.Info("starting app")

	api.Run(a.cfg)
	a.logger.Info("database starting")
	database, err := db.New(*a.cfg)

	if err != nil {
		return err
	}

	err = database.AutoMigrate()
	if err != nil {
		return err
	}

	a.logger.Info("bot starting")
	_ = db.NewService(database)

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("stopping app")
	return nil
}
