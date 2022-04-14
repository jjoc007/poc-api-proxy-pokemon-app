package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		panic(err)
	}
}

func run(port string) error {
	router := gin.Default()
	health := HealthChecker{}
	mapRoutes(router, health)
	return router.Run(":" + port)
}

func mapRoutes(r *gin.Engine, health HealthChecker) {
	r.GET("/ping", health.PingHandler)
	//router.URLMappings(r, router.Build())
}

type HealthChecker struct{}

func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}