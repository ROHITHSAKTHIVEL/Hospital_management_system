package config

import (
	"context"
	"ecommerce/constants"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Customer_Collection *mongo.Collection
var Customer_ProfileCollection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Customer_ProfileCollection = client.Database(constants.DB_Name).Collection(constants.Customer_collection)

	fmt.Println("Collection Connected")
}

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Customer_Collection = client.Database(constants.DB_Name).Collection(constants.Customer_appoitment)

	fmt.Println("Collection Connected")
}
