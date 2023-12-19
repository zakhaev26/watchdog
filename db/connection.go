package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The specific collection instance declared in the global namespace
var Collection *mongo.Collection

// Initiator function which runs automatically at the start of application
// This function connects to the DB and returns collection instances which is used for DB actions.
func Init() {
	const connectionString = "mongodb+srv://linux-skg:1TuX01zH2y3tjUFV@sports.vj9j4tb.mongodb.net/?retryWrites=true&w=majoritymongodb+srv://linux-skg:1TuX01zH2y3tjUFV@sports.vj9j4tb.mongodb.net/?retryWrites=true&w=majority"
	const databaseName = "sse_part2"
	const collection1_Name = "test_v1"

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	//handling error
	if err != nil {
		log.Fatal("Error Occurred while connecting to the database:", err)
	}

	//if passes through this,connection is successful

	fmt.Println("Connection to Database Successful")

	//assigning collection
	Collection = client.Database(databaseName).Collection(collection1_Name)
	fmt.Printf("Collection Instance %s is Ready.", collection1_Name)
}