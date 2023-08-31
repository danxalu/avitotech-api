package main

import (
	"avitotech-api/cfg"
	"github.com/golang-migrate/migrate/v4"
	"fmt"
	"os"
)

func main() {
	conf, err := cfg.ApplyConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DBName)
	migration, err := migrate.New("file:///migrations", url)

	if err != nil { fmt.Println(err) }
	if req := migration.Up(); req != nil { fmt.Println(err)}
}
