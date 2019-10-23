package main

import (
	dbConn "bass/adapter/gorm"
	"bass/config"
	"bass/server/app"
	"bass/server/router"
	lr "bass/util/logger"
	vr "bass/util/validator"
	"fmt"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}

	if appConf.Debug {
		db.LogMode(true)
	}

	validator := vr.New()

	application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
