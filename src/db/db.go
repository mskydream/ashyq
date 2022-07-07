package db

import (
	"fmt"

	"github.com/mskydream/qr-code/src/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Conn *sqlx.DB
}

func (d *DB) InitDatabase(c *config.DB) *DB {
	var err error
	source := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Address, c.Name)

	if d.Conn, err = sqlx.Connect("pgx", source); err != nil {
		panic(err)
	}

	m, err := migrate.New("file://migrations", source)
	if err != nil {
		panic(err)
	}
	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}
	return d
}
