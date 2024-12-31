package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sundayonah/phindcode_backend/internal/config"
	"github.com/sundayonah/phindcode_backend/internal/handlers"
	"github.com/sundayonah/phindcode_backend/internal/service"
	"github.com/sundayonah/phindcode_backend/pkg/middleware"
)

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create config
	cfg := config.NewConfig()
	defer cfg.CloseDB()

	svc := service.NewPostService(cfg.Client)
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
