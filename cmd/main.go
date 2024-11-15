package main

import (
	"log"
	"os"
	"task-doodocs/app"
	"task-doodocs/internal/config"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conf := config.Loader()
	app := app.New(infoLog, errorLog)

	repo, err := repo.New(conf.StoragePath)
	if err != nil {
		errorLog.Fatal(err)
	}
}
