package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type DataFlightUnProcess struct {
	Origin      interface{}
	Destination interface{}
}

type DataFlight struct {
	Origin      AirPortDetail
	Destination AirPortDetail
}

type AirPortDetail struct {
	TimeZone interface{}
	Name string
	Code interface{}
	Position Position
}

type Position struct {
	Latitude float64
	Longitude float64
	Country interface{}
	Region Region
}

type Region struct {
	City string
}

func ConnectMongo() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to MongoDB Success!")
	}
	return client
}

func SaveDataCrawlerToMongo(client *mongo.Client, data []DataFlightUnProcess)  {
	collection := client.Database("jx_test").Collection("JxTest")
	count := 0
	for _ , v := range data {
		_, err := collection.InsertOne(context.TODO(), bson.D{
			{"origin", v.Origin},
			{"destination", v.Destination},
		})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			count ++
		}
	}
	fmt.Println("Nums item to MongoDB from Fight24Radar: ", count)
}

func GetAllFlyJV(client *mongo.Client)  (list []DataFlight, err error){
	collection := client.Database("jx_test").Collection("JxTest")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	// Find() method raised an error
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			var result DataFlight
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				list = append(list, result)
			}
		}
	}
	return
}

func DisconnectMongo(client *mongo.Client)  {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
