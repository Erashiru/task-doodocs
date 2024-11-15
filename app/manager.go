package app

import "log"

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func New(infoLog, errorLog *log.Logger) *Application {
	return &Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
