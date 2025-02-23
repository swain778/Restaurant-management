package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_item", controllers.GetOrderItems())
	incomingRoutes.GET("/order_item/:order_id", controllers.GetOrderItem())
	incomingRoutes.POST("/order_item", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order_item/:order_id", controllers.UpdateOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
}
