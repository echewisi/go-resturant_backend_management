package routes

import (
	controllers "resturant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func foodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.getFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.getFood())
	incomingRoutes.POST("/create_food", controllers.createFood())
	incomingRoutes.PATCH("/update_food/:food_id", controllers.updateFood())
}
