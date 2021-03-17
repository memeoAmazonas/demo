package persistence

import (
	"../log"
	"context"
	"fmt"

	c "../model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

const FILE_NAME = "config.go"

func init() {
	clientOptions := options.Client().ApplyURI(c.URL_BD)
	client, err := mongo.Connect(ctx, clientOptions)
	var logResponse c.LogResponse
	logResponse.FileLocation = FILE_NAME
	if err != nil {
		logResponse.Line = "20"
		logResponse.Message = c.CONNECTION_DB_ERROR
		log.LogError(logResponse)
	}
	//noinspection ALL
	err = client.Ping(ctx, nil)
	if err != nil {
		logResponse.Line = "27"
		logResponse.Message = c.CONNECTION_DB_ERROR_PING
		log.LogError(logResponse)
	}
	collection = client.Database(c.DB_NAME).Collection(c.COLLECTION)
	fmt.Println(c.CONNECTION_DB_SUCCESSFULL)
}
