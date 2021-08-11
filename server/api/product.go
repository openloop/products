package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/openloop/products/server/env"
	"github.com/openloop/products/server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request, e *env.Env) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	cur, err := e.MongoCollection.Find(r.Context(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(r.Context())

	var products []models.Product

	for cur.Next(r.Context()) {
		var product models.Product
		if err = cur.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	j := json.NewEncoder(w)
	j.SetIndent("", "  ")
	j.Encode(products)

	/*
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(e.Products)
		w.WriteHeader(http.StatusOK)
	*/
}

func GetProductById(w http.ResponseWriter, r *http.Request, e *env.Env) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idParam, idParamExists := vars["productId"]

	idInt, idErr := strconv.ParseInt(idParam, 10, 64)
	filter := bson.M{"id": idInt}
	var product models.Product

	mErr := e.MongoCollection.FindOne(r.Context(), filter).Decode(&product)

	if idParamExists && idErr == nil && mErr == nil {
		j := json.NewEncoder(w)
		j.SetIndent("", "  ")
		j.Encode(product)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println(mErr)
		http.Error(w, http.StatusText(http.StatusBadRequest)+": Invalid ProductID", http.StatusBadRequest)
	}
	/*
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		vars := mux.Vars(r)
		p, exists := vars["productId"]
		productID := new(big.Int)
		productID, converted := productID.SetString(p, 10)

		if exists && converted {
			for i, product := range e.Products {
				if product.ID.Cmp(productID) == 0 {
					json.NewEncoder(w).Encode(e.Products[i])
					w.WriteHeader(http.StatusOK)
					return
				}
			}
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest)+": Invalid ProductID", http.StatusBadRequest)
			return
		}

		http.Error(w, http.StatusText(http.StatusNotFound)+": ProductID not found", http.StatusNotFound)
	*/
}
