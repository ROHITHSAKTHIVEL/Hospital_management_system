package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Getalldata() []primitive.M {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []primitive.M
	for cursor.Next(context.Background()) {
		var profile bson.M
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}

func Insert(profile models.Customer) {

	inserted, err := config.Customer_ProfileCollection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", inserted.InsertedID)
}

func Appoitment(profile models.Appoitment) {
	inserted, err := config.Customer_Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", inserted.InsertedID)
}
func Login(profile models.Login) error {
	fmt.Println("service")
	var ctx context.Context
  //  var login models.Login
	query := bson.M{"email": profile.Email, "password": profile.Password}
	var customer models.Customer
	err := config.Customer_ProfileCollection.FindOne(ctx,query).Decode(&customer)
	if err != nil {
		fmt.Println(err)
	}
	return err
	
}
