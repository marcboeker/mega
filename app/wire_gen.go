// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/db"
	"github.com/marcboeker/mega/logger"
	"github.com/marcboeker/mega/server"
	"github.com/marcboeker/mega/service"
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
	mapper := service.New(client, logLogger)
	serverServer := server.New(configConfig, mapper, logLogger)
	app := New(configConfig, client, serverServer)
	return app, nil
}
