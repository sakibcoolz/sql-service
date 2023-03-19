package utils

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func ResponseWriter(w http.ResponseWriter, statusCode int, response interface{}, log *zap.Logger) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Error("Received error on marshal response",
			zap.Error(err))

		statusCode = http.StatusUnprocessableEntity
		jsonResponse = []byte(err.Error())
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonResponse); err != nil {
		log.Error("Error while http response", zap.Error(err))
	}
}
