package loader

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/openloop/products/server/models"
)

const PRODUCTS_URL = "https://my-json-server.typicode.com/convictional/engineering-interview/products"

func LoadProducts() []models.Product {
	resp, err := http.Get(PRODUCTS_URL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	productJSON, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var products []models.Product
	json.Unmarshal([]byte(productJSON), &products)
	return products
}
