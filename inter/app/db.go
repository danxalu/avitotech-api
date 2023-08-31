package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (app *Application) connectDB() (*sql.DB, error) {
	url_db := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		app.config.Postgres.Host, app.config.Postgres.Port,
		app.config.Postgres.User, app.config.Postgres.Password,
		app.config.Postgres.DBName)

	db, err := sql.Open("db", url_db)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
