package routes

import (
	"github.com/gin-gonic/gin"
	controllers "resturant_backend/controllers"

)

func foodRoutes( incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controllers.getFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.getFood())
	incomingRoutes.POST("/create_food", controllers.createFood())
	incomingRoutes.PATCH("/update_food/:food_id", controllers.updateFood())
}