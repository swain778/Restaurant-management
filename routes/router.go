package routes

import (
	"restaurant-management/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/foods", handlers.CreateFood)
	r.GET("/foods/:food_id", handlers.GetFoodById)
	r.GET("/foods", handlers.GetFoods)
	r.DELETE("/foods/:food_id", handlers.DeleteFood)

	r.POST("/invoice", handlers.CreateInvoice)
	r.GET("/invoice/:invoice_id", handlers.GetInvoiceById)
	r.GET("/invoices", handlers.GetInvoices)
	r.DELETE("/invoice/:invoice_id", handlers.DeleteInvoice)

	r.POST("/menu", handlers.CreateMenu)
	r.GET("/menu/:menu_id", handlers.GetMenuById)
	r.GET("/menus", handlers.GetMenus)
	r.DELETE("/menu/:menu_id", handlers.DeleteMenu)

	r.POST("/order", handlers.CreateOrder)
	r.GET("/order/:order_id", handlers.GetOrderById)
	r.GET("/orders", handlers.GetOrders)
	r.DELETE("/orders/:invoice_id", handlers.DeleteOrder)

	r.POST("/order_item", handlers.CreateOrderItem)
	r.GET("/order_item/:order_item_id", handlers.GetOrderItemById)
	r.GET("/order_items", handlers.OrderItems)
	r.DELETE("/order_item/:order_item_id", handlers.DeleteOrderItem)

	return r
}
