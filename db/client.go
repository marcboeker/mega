package db

import (
	"context"
	"log"

	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/ent"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	*ent.Client
}

func New(cfg *config.Config) (*Client, error) {
	var opts []ent.Option
	if cfg.GetBool("db.debugging") {
		opts = append(opts, ent.Debug())
	}
	cli, err := ent.Open(cfg.GetString("db.driverName"), cfg.GetString("db.dataSourceName"), opts...)
	if err != nil {
		return nil, err
	}
	if err := cli.Schema.Create(context.Background()); err != nil {
		log.Fatalf("error: failed creating schema resources %v\n", err)
	}
	return &Client{cli}, err
}

func (db *Client) Close() error {
	return db.Client.Close()
}
