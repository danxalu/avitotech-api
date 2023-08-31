package app

import (
	"avitotech-api/cfg"
	"avitotech-api/inter/cases"
	"avitotech-api/inter/database"
	"avitotech-api/inter/server"
	"fmt"
	"os"
)

type Application struct {
	config *cfg.Config
}

func (app *Application) Start() error {
	db, err := app.connectDB()
	if err != nil {
		return err
	}

	repositories := database.Rep_New(db)
	myUseCase := cases.NewUC(repositories)
	newServer := server.NewServer(*app.config, myUseCase)

	err = newServer.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
	return nil
}