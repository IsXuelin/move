package main

import (
	"github.com/gin-gonic/gin"
	"move/controller"
)

//
//func CollectRoute(router *gin.Engine) *gin.Engine {
//	orderOpe := router.Group("/orders")
//	{
//		orderOpe.POST("", controller.PlaceOrder)
//		orderOpe.PATCH("/:id", controller.TakeOrder)
//		orderOpe.GET("", controller.OrderList)
//	}
//	return router
//}

func CollectRoute(router *gin.Engine) *gin.Engine {

	router.POST("/orders", controller.PlaceOrder)
	router.PATCH("/orders/:id", controller.TakeOrder)
	router.GET("/orders", controller.OrderList)

	return router
}
