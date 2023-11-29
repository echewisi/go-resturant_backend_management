package routes

import (
	"github.com/gin-gonic/gin"
	controllers "resturant_backend/controllers"
)

func invoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("invoice/", controllers.getInvoices())
	incomingRoutes.GET("invoice/:invoice_id", controllers.getInvoice())
	incomingRoutes
}