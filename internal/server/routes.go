package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	// Use middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS setup
	allowOrigins := strings.Split(os.Getenv("ALLOW_ORIGIN"), ",")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowOrigins,
	}))

	// Default group
	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.healthHandler)

	// Example group
	exampleGroup := e.Group("/examples")
	exampleGroup.GET("", s.listExample)    // GET /examples
	exampleGroup.GET("/:id", s.getExample) // GET /examples/:id
	exampleGroup.POST("", s.addExample)    // POST /examples

	return e
}
