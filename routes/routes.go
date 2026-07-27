package routes

import (
	"GO/SubjectGIN_API_Food_Delivery_Service/handlers"

	"github.com/gin-gonic/gin"
)

func GetOrder(r *gin.Engine) {
	r.GET("/orders", handlers.GetOrdersHandler)
}

func GetOrderId(r *gin.Engine) {
	r.GET("/orders/:id", handlers.GetOrdersIdHandler)
}

func GetOrdersSearch(r *gin.Engine) {
	r.GET("/orders/search?status=cooking", handlers.GetOrdersSearchHandler)
}

func PostOrders(r *gin.Engine) {
	r.POST("/orders", handlers.PostOrdersHandler)
}

func PatchOrders(r *gin.Engine) {
	r.PATCH("/orders/:id/status/:new_status", handlers.PatchOrderHandler)
}
