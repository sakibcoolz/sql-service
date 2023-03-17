package app

import (
	"encoding/json"
	"net/http"
	"sql-service/config"
	"sql-service/database"
	"time"

	"github.com/gorilla/mux"
)

func MapURL(configSetup *config.Config, db *database.DBInstance) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	return &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
