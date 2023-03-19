package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sql-service/config"
	"sql-service/contoller"
	"sql-service/database"
	"sql-service/domain"
	"sql-service/service"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func MapURL(configSetup *config.Config, db *database.DBInstance) *http.Server {
	go TerminateService(StopService, configSetup.Log)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

	controller := &contoller.Controller{
		Service: &service.Service{
			Storage: domain.NewStorage(db.DB, configSetup.Log),
			Log:     configSetup.Log,
		},
		Log: configSetup.Log,
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", controller.Health).Methods("GET")

	router.HandleFunc("/sqlconsole", controller.SQLConsole).Methods("GET")

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%d", configSetup.Service.Host, configSetup.Service.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	select {
	case <-stopService:
		log.Info("Terminating service")

		os.Exit(0)
	}
}
