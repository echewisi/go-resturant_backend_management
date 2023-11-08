package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controllers "resturant_backend/controllers"

)

func userRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controllers.getUsers())
	incomingRoutes.GET("/users/:user_id", controllers.getUser())
	incomingRoutes.POST("/signup", controllers.signupUser())
	incomingRoutes.POST("/login", controllers.loginUser())

}