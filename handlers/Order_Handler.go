package handlers

import (
	"GO/SubjectGIN_API_Food_Delivery_Service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrdersHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Orders)

}

func GetOrdersIdHandler(c *gin.Context) {
	id := c.Param("id")

	order, ok := models.Orders[id]

	if ok {
		c.IndentedJSON(http.StatusOK, order)
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

}

func GetOrdersSearchHandler(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Параметр status обязателен"})
	}

	c.IndentedJSON(http.StatusOK, models.Orders)
}

func PostOrdersHandler(c *gin.Context) {
	var NewOrder models.NewOrder

	if err := c.ShouldBindJSON(&NewOrder); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if NewOrder.CustomerName == "" || NewOrder.Items == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Имя должно быть заполнено и корзина должна быть не пуста"})
		return
	}
	newId := strconv.Itoa(len(models.Orders) + 1)
	total := 0.0
	for _, v := range NewOrder.Items {

		total += v.Price * float64(v.Quantity)
	}
	order := &models.Order{
		ID:           newId,
		CustomerName: NewOrder.CustomerName,
		Status:       "Created",
		Items:        NewOrder.Items,
		TotalPrice:   total,
	}

	models.Orders[newId] = order

	c.IndentedJSON(http.StatusCreated, order)

}

func PatchOrderHandler(c *gin.Context) {
	id := c.Param("id")
	newStatus := c.Param("new_status")

	order, ok := models.Orders[id]

	if ok {
		if newStatus == "cooking" || newStatus == "delivered" || newStatus == "cancelled" {
			if newStatus == "delivered" || newStatus == "cancelled" {
				c.IndentedJSON(http.StatusConflict, gin.H{"error": "Нельзя изменить статус уже завершенного заказа"})
				return
			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Недопустимый статус"})
				return
			}
		}
		c.IndentedJSON(http.StatusOK, order.Status == newStatus)
		return
	}

}
