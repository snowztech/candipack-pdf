package main

import (
	"candipack-pdf/configs"
	"candipack-pdf/internal/handlers"
	"candipack-pdf/internal/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.Load()
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORS())
	if config.APIKey != "" {
		router.Use(middleware.APIKey(config.APIKey))
	}

	// Initialize handlers
	h := handlers.New()

	apirouter := router.Group("/api")

	// Routes
	apirouter.POST("/resume", h.HandleResume())
	apirouter.POST("/resume/html", h.HandleResumeHTML())
	apirouter.POST("/cover-letter", h.HandleCoverLetter())
	router.GET("/templates", h.HandleTemplates())
	router.GET("/up", func(c *gin.Context) {
		c.String(200, "ok")
	})

	log.Printf("Server running on :%d", config.Port)
	log.Fatal(router.Run(":" + fmt.Sprint(config.Port)))
}
