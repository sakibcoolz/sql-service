package main

import (
	"sql-service/app"
	"sql-service/config"
	"sql-service/database"

	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

func main() {
	configSetup := config.NewConfigSetup(log, "sql-manager")

	configSetup.GetConfiguration()

	if err := app.MapURL(configSetup, database.Connection(configSetup)).ListenAndServe(); err != nil {
		log.Fatal("Service terminated", zap.Error(err))
	}
}
