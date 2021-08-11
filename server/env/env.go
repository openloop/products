package env

import (
	"github.com/openloop/products/server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_URI        = "mongodb://localhost:27017"
	MONGO_DB         = "products"
	MONGO_COLLECTION = "products"
)

type Env struct {
	Products        []models.Product
	MongoCollection *mongo.Collection
}
