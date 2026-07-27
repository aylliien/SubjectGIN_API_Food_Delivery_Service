package models

type Order struct {
	ID           string      `json:"id"`
	CustomerName string      `json:"customer_name"`
	Status       string      `json:"status"`
	Items        []OrderItem `json:"items"`
	TotalPrice   float64     `json:"total_price"`
}

type NewOrder struct {
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
}

var Orders = map[string]*Order{
	"1": {
		ID: "1", CustomerName: "Billy Wilamson", Status: "Created", Items: []OrderItem{{"Pizza", 3200, 1}}, TotalPrice: 3200,
	},
	"2": {
		ID: "2", CustomerName: "Dick Medown", Status: "Cooking", Items: []OrderItem{{"Khan Burger", 4500, 1}, {"Coke 1L", 1200, 1}, {"Chips", 950, 1}}, TotalPrice: 6650,
	},
}
