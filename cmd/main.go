package main

import (
	"os/signal"
	"sql-service/app"
	"sql-service/config"
	"sql-service/database"
	"syscall"

	"go.uber.org/zap"
)

func main() {

	log := zap.NewExample()

	configSetup := config.NewConfigSetup(log, "sql-manager")

	configSetup.GetConfiguration()

	if err := app.MapURL(configSetup, database.Connection(configSetup)).ListenAndServe(); err != nil {
		log.Error("Service terminated", zap.Error(err))
		signal.Notify(app.StopService, syscall.SIGINT, syscall.SIGTERM)
	}
}
