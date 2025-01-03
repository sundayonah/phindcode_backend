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
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Create config
	client, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	defer config.CloseDB()

	// Set up services
	svc := service.NewPostService(client)
	handler := handlers.NewPostHandler(svc)

	// Set up Auth service and handler
	authSvc := service.NewAuthService(client)
	authHandler := handlers.NewAuthHandler(authSvc)

	// Initialize the Gin router
	r := gin.Default()

	// Apply the auth middleware to all routes

	// Social service setup
	socialSvc := service.NewSocialService(client)
	socialHandler := handlers.NewSocialHandler(socialSvc)

	// Define version 1 API group
	v1 := r.Group("/api/v1")
	{
		// Register the /login route for user login
		v1.GET("/users", authHandler.GetAllUsers)
		v1.POST("/login", authHandler.LogIn)
		v1.POST("/register", authHandler.Register)

		// Apply middleware to other routes
		authorized := v1.Group("/")
		authorized.Use(middleware.AuthMiddleware()) // Apply the auth middleware here for protected routes

		// Post routes
		posts := authorized.Group("/posts")

		{
			posts.POST("/", handler.CreatePost)
			posts.GET("/", handler.GetPosts)
			posts.GET("/:id", handler.GetPost)
			posts.PUT("/:id", handler.UpdatePost)
			posts.DELETE("/:id", handler.DeletePost)

			// Social routes
			handlers.AddSocialRoutes(posts, socialHandler)
		}
	}

	// Define the port and start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run(":" + port))
}
