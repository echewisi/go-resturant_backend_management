package routes

import (
	controllers "resturant_backend/controllers"
	"github.com/gin-gonic/gin"
)

func tableRoutes(incomingroutes *gin.Engine){
	incomingroutes.GET("/gettables", controllers.getTables())
	incomingroutes.GET("/gettable/:table_id", controllers.getTable())
}