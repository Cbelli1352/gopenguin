package controllers

import (
	"database/sql"

	"gopenguin/auth"
	"gopenguin/config"

	_ "github.com/mattn/go-sqlite3"
)

type Controller struct {
	DB   *sql.DB
	Auth *auth.Auth
}

func NewController(config *config.Config) (*Controller, error) {
	var err error

	c := new(Controller)

	c.Auth, err = auth.NewAuth(config)

	if err != nil {
		return nil, err
	}

	c.DB, err = sql.Open("sqlite3", config.Database.Source)

	return c, err
}
