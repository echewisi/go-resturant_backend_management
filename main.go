package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"resturant_backend/routes"
	"go.mongo.org/mongo-driver/mongo"
)

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port= "8000"
	}

	router:= gin.Now()
	router.Use(gin.Logger())
}