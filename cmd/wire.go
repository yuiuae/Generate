package main

import (
	"database/sql"
	"log"

	"github.com/google/wire"
)

type Container struct {
	Logger *log.Logger
	DB     *sql.DB
}

// wire - considered as an example without real use in the application
// (did not find where to apply:)

func NewContainer() (*Container, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Container), "*"),
	))
}
