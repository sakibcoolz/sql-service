package app

import (
	"fmt"
	"net/http"
	"sql-service/config"
	"sql-service/contoller"
	"sql-service/database"
	"sql-service/domain"
	"sql-service/service"
	"time"

	"github.com/gorilla/mux"
)

func MapURL(configSetup *config.Config, db *database.DBInstance) *http.Server {

	controller := &contoller.Controller{
		Service: &service.Service{
			Storage: domain.NewStorage(db.DB, configSetup.Log),
			Log:     configSetup.Log,
		},
		Log: configSetup.Log,
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", controller.Health).Methods("GET")

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%d", configSetup.Service.Host, configSetup.Service.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
