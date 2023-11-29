package routes

import (
	controllers "resturant_backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes( incomingroutes *gin.Engine){
	incomingroutes.GET("/get_orders", controllers.GetOrders())
	incomingroutes.GET("/get_orders/:order_id", controllers.GetOrder())
	incomingroutes.POST("/create_order", controllers.GetOrder())
	incomingroutes.PATCH("/update_order/:order_id", controllers.updateOrder())
}