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

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port= "8000"
	}

	router:= gin.New()
	router.Use(gin.Logger())
	routes.userRouter(router)
	router.Use(middleware.Authentication)

	routes.foodRouter(router)
	routes.menuRouter(router)
	routes.tableRouter(router)
	routes.orderRouter(router)
	routes.orderItemRouter(router)
	routes.invoiceRouter(router)
}