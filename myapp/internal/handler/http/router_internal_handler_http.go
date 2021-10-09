package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(customerHandler *CustomerHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.GET("/", Status)
	engine.GET("/health", Health)
	engine.GET("/version", Version)

	engine.POST("/create-customer", customerHandler.CreateCustomer)
	engine.GET("/list-customer", customerHandler.GetListCustomer)
	engine.GET("/get-customer/:id", customerHandler.GetDetailCustomer)
	engine.PUT("/update-customer/:id", customerHandler.UpdateCustomer)
	engine.DELETE("/delete-customer/:id", customerHandler.DeleteCustomer)

	return engine
}
