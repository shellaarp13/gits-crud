package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(internalUsername, internalPassword string) *echo.Echo {
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

	// engine.POST("/create-entity", entityHandler.CreateEntity)
	// engine.GET("/list-entity", entityHandler.GetListEntity)
	// engine.GET("/get-entity/:id", entityHandler.GetDetailEntity)
	// engine.PUT("/update-entity/:id", entityHandler.UpdateEntity)
	// engine.DELETE("/delete-entity/:id", entityHandler.DeleteEntity)

	return engine
}
