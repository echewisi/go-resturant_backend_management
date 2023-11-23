package routes

import (
	controllers "resturant_backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetItemsOrder_order())
	incomingRoutes.POST("/orderItems/:orderItem_id",  controllers.CreateOrderItem())
}