package api

import (
	"net/http"

	"github.com/openloop/products/server/env"
)

func GetInventory(w http.ResponseWriter, r *http.Request, e *env.Env) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
