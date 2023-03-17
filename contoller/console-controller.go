package contoller

import (
	"encoding/json"
	"net/http"
	"sql-service/service"

	"go.uber.org/zap"
)

type Controller struct {
	Service service.IService
	Log     *zap.Logger
}

func (c *Controller) Console() {
	c.Service.Console()
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {

	c.Log.Info("Health API Called")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
