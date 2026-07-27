package main

import (
	"GO/SubjectGIN_API_Food_Delivery_Service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.GetOrder(r)

	routes.GetOrderId(r)

	routes.GetOrdersSearch(r)

	routes.PostOrders(r)
}
