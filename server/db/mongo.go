package db

import (
	"context"
	"fmt"
	"log"

	"github.com/openloop/products/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI        string
	DB         string
	Collection string
}

func (m MongoConfig) GetMongoConnection() (*mongo.Client, error) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))

	if err != nil {
		log.Println(err)
	}

	return client, nil
}

func (m MongoConfig) GetMongoCollection() *mongo.Collection {
	client, err := m.GetMongoConnection()

	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to connect to MongoDB: %s", err))
	}

	collection := client.Database(m.DB).Collection(m.Collection)

	return collection
}

func (m MongoConfig) GetAllProducts(c context.Context) []models.Product {
	cur, err := m.GetMongoCollection().Find(c, bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(c)

	var products []models.Product

	for cur.Next(c) {
		var product models.Product
		if err = cur.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}
