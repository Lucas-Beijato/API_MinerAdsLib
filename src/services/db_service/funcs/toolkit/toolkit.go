package dbtoolkit

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Test connection mongo
func Ping_Connection_DataBase(client *mongo.Client, context *context.Context) error {
	err := client.Database("admin").RunCommand(*context, bson.D{{Key: "ping", Value: 1}}).Err()
	// Send a ping to confirm a successful connection
	if err != nil {
		return err
	}
	return nil
}

// Create a New Client Mongo
func Gen_ServerAPIClient(context *context.Context) (*mongo.Client, error) {
	con_String, is_present_con_string := os.LookupEnv("CONN_STRING_DB_MONGO")
	if !is_present_con_string {
		fmt.Println("Invalid connection DataBase String.")
		return nil, errors.New("error to obtain the connection string")
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(con_String).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(*context, opts)
	if err != nil {
		return nil, err
	}

	return client, err
}

// Close connection mongo
func Close_DataBase_Connection(client *mongo.Client, context *context.Context) error {

	err := client.Disconnect(*context)
	if err != nil {
		return err
	}
	return nil
}
