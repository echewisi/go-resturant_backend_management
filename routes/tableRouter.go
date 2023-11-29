package routes

import (
	controllers "resturant_backend/controllers"
	"github.com/gin-gonic/gin"
)

func tableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/get_tables", controllers.getTables())
	incomingRoutes.GET("/get_table/:table_id", controllers.getTable())
	incomingRoutes.POST("/create_table", controllers.createTable())
	incomingRoutes.PUT("/update_table/:table_id", controllers.updateTable())
}