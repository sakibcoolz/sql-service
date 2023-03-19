package contoller

import (
	"encoding/json"
	"net/http"
	"sql-service/model"
	"sql-service/service"
	"sql-service/utils"

	"go.uber.org/zap"
)

type Controller struct {
	Service service.IService
	Log     *zap.Logger
}

func (c *Controller) SQLConsole(w http.ResponseWriter, r *http.Request) {
	var request model.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Error("Request body incorrect.", zap.Error(err))

		utils.ResponseWriter(w, http.StatusBadRequest, struct {
			name string `json:"name"`
		}{name: "Bad Request"}, c.Log)

		return
	}

	response := c.Service.Console(request)

	utils.ResponseWriter(w, http.StatusOK, response, c.Log)

	return
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {

	c.Log.Info("Health API Called")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
