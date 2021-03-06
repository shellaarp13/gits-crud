package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(accountHandler *AccountHandler, customerHandler *CustomerHandler, orderHandler *OrderHandler, productHandler *ProductHandler, orderdetailsHandler *Order_DetailsHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	engine.GET("/", Status)
	engine.GET("/health", Health)
	engine.GET("/version", Version)
	engine.POST("/register", accountHandler.CreateAccount)

	engine.POST("/create-account", accountHandler.CreateAccount)
	engine.GET("/list-account", accountHandler.GetListAccount)
	engine.GET("/get-account/:id", accountHandler.GetDetailAccount)
	engine.PUT("/update-account/:id", accountHandler.UpdateAccount)
	engine.DELETE("/delete-account/:id", accountHandler.DeleteAccount)

	engine.POST("/create-customer", customerHandler.CreateCustomer)
	engine.GET("/list-customer", customerHandler.GetListCustomer)
	engine.GET("/get-customer/:id", customerHandler.GetDetailCustomer)
	engine.PUT("/update-customer/:id", customerHandler.UpdateCustomer)
	engine.DELETE("/delete-customer/:id", customerHandler.DeleteCustomer)

	engine.POST("/create-order", orderHandler.CreateOrder)
	engine.GET("/list-order", orderHandler.GetListOrder)
	engine.GET("/get-order/:id", orderHandler.GetDetailOrder)
	engine.PUT("/update-order/:id", orderHandler.UpdateOrder)
	engine.DELETE("/delete-order/:id", orderHandler.DeleteOrder)

	engine.POST("/create-product", productHandler.CreateProduct)
	engine.GET("/list-product", productHandler.GetListProduct)
	engine.GET("/get-product/:id", productHandler.GetDetailProduct)
	engine.PUT("/update-product/:id", productHandler.UpdateProduct)
	engine.DELETE("/delete-product/:id", productHandler.DeleteProduct)

	engine.POST("/create-orderdetails", orderdetailsHandler.CreateOrder_Details)
	engine.GET("/list-orderdetails", orderdetailsHandler.GetListOrder_Details)
	engine.GET("/get-orderdetails/:id", orderdetailsHandler.GetDetailOrder_Details)
	engine.PUT("/update-orderdetails/:id", orderdetailsHandler.UpdateOrder_Details)
	engine.DELETE("/delete-orderdetails/:id", orderdetailsHandler.DeleteOrder_Details)

	return engine
}
