package routes

import (
	"github.com/gin-gonic/gin"
	controllers "resturant_backend/controllers"
)

func menuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menu", controllers.getMenus())
	incomingRoutes.GET("/menu/:menu_id", controllers.getMenu())
	incomingRoutes.POST("/createMenu", controllers.createMenu())
	incomingRoutes.PATCH("/updateMenu/:menu_id", controllers.updateMenu())
}