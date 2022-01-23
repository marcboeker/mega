// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package test

import (
	"context"
	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/db"
	"github.com/marcboeker/mega/logger"
	"log"
)

// Injectors from wire.go:

func Initialize() (*App, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	client, err := db.New(configConfig)
	if err != nil {
		return nil, err
	}
	logLogger := logger.New()
	app := New(configConfig, client, logLogger)
	return app, nil
}

// wire.go:

type App struct {
	Cfg    *config.Config
	Client *db.Client
	Ctx    context.Context
	Logger *log.Logger
}

func New(cfg *config.Config, client *db.Client, logger2 *log.Logger) *App {
	return &App{cfg, client, context.Background(), logger2}
}

func (a *App) Close() error {
	return a.Client.Close()
}
