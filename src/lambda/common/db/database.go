package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"invoice/common"
	"log"
	"os"
)

func CloseDbConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	common.LogError("Error when disconnecting the database", err)
}

func OpenDbConnection() *mongo.Client {
	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatal("It is missing database URI")
	}
	opts := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		common.LogError("Error when connecting the database", err)
	}
	return client
}

func GetInvoiceCollection(conn *mongo.Client) *mongo.Collection {
	return conn.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))
}
