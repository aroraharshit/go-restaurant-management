package routes

import (
	controller "go-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoives/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("invoices/:invoice_id", controller.UpdateInvoice())
}
