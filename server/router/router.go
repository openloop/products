package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openloop/products/server/api"
	"github.com/openloop/products/server/env"
	"github.com/openloop/products/server/logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(e *env.Env) *mux.Router {

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},

		Route{
			"GetAllProducts",
			"GET",
			"/products",
			func(w http.ResponseWriter, r *http.Request) {
				api.GetAllProducts(w, r, e)
			},
		},

		Route{
			"GetProductById",
			"GET",
			"/products/{productId}",
			func(w http.ResponseWriter, r *http.Request) {
				api.GetProductById(w, r, e)
			},
		},

		Route{
			"GetInventory",
			"GET",
			"/store/inventory",
			func(w http.ResponseWriter, r *http.Request) {
				api.GetInventory(w, r, e)
			},
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product API")
}
