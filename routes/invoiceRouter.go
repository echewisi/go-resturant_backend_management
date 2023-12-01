package routes

import (
	controllers "resturant_backend/controllers"
	"github.com/gin-gonic/gin"
)

func invoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("invoice/", controllers.getInvoices())
	incomingRoutes.GET("invoice/:invoice_id", controllers.getInvoice())
	incomingRoutes.POST("/createInvoice", controllers.createInvoice())
	incomingRoutes.PATCH("/updateInvoice/:invoice_id", controllers.updateInvoice())
}
