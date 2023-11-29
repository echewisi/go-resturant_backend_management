package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println(client)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second){
	
	defer cancel()

	client.Connect(ctx)

	if err != nil{
			log.Fatal(err)
		} else{
			fmt.Println("connected to mongodb database")
			return client
		}
}