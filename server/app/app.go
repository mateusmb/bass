package app

import (
	"bass/util/logger"

	"github.com/jinzhu/gorm"
)

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{logger: logger, db: db}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
