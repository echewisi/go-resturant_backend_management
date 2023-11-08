package main

import (
	"github.com/gin-gonic/gin"
	"os"
	routes "resturant_backend/routes"
	middleware "resturant_backend/middleware"
	controllers "resturant_backend/controllers"
	database "resturant_backend/database"
	"go.mongo.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection= database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port= "8000"
	}

	router:= gin.New()
	router.Use(gin.Logger())
	routes.userRoutes(router)
	router.Use(middleware.Authentication())

//engaging the routes
	routes.foodRoutes(router)
	routes.menuRoutes(router)
	routes.tableRoutes(router)
	routes.orderRoutes(router)
	routes.orderItemRoutes(router)
	routes.invoiceRoutes(router)

	router.Run(":" + port)
}