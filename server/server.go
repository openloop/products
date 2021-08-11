package main

import (
	"log"
	"net/http"

	"github.com/openloop/products/server/db"
	"github.com/openloop/products/server/env"
	"github.com/openloop/products/server/loader"
	"github.com/openloop/products/server/router"
)

func main() {
	log.Printf("Server started")

	// Read some sample product data
	products := loader.LoadProducts()

	// Create shared enviroment

	d := db.MongoConfig{URI: env.MONGO_URI, DB: env.MONGO_DB, Collection: env.MONGO_COLLECTION}

	e := env.Env{Products: products, MongoCollection: d.GetMongoCollection()}

	// Create a router for http requests
	router := router.NewRouter(&e)

	// Serve requests
	log.Fatalln(http.ListenAndServe(":8080", router))
}
