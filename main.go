package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sundayonah/phindcode_backend/internal/config"
	"github.com/sundayonah/phindcode_backend/internal/handlers"
	"github.com/sundayonah/phindcode_backend/internal/service"
	"github.com/sundayonah/phindcode_backend/pkg/middleware"
)

func main() {
	// load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Create config
	client, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	defer config.CloseDB()

	svc := service.NewPostService(client)
	handler := handlers.NewPostHandler(svc)

	r := gin.Default()

	// Middleware
	r.Use(middleware.AuthMiddleware())

	// Routes
	v1 := r.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.POST("/", handler.CreatePost)
			posts.GET("/", handler.GetPosts)
			posts.GET("/:id", handler.GetPost)
			posts.PUT("/:id", handler.UpdatePost)
			posts.DELETE("/:id", handler.DeletePost)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run(":" + port))
}
