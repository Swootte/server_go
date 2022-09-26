package database

import (
	"context"
	"fmt"
	"os"

	"server/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	utils.LoadEnv()
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@cluster0.myzbc.mongodb.net/"+os.Getenv("DATABASE")+"?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	MongoClient = client

	// defer func() {
	// 	if err = MongoClient.Disconnect(ctx); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
}
