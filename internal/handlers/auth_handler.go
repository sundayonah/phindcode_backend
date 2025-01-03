package handlers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sundayonah/phindcode_backend/internal/service"
	"golang.org/x/crypto/bcrypt"
)

// SignInRequest defines the login request structure
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"` // For Google login
	Method   string `json:"method"`          // "email" or "google"
}

// LogIn handles user login for both email/password and Google login
func LogIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	switch req.Method {
	case "email":
		// Handle email/password login
		token, err := handleEmailLogin(c, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

	case "google":
		// Handle Google login (implement the Google verification as you did earlier)
		token, err := handleGoogleLogin(c, req.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login method"})
	}
}

// handleEmailLogin validates the user's credentials and generates a JWT token
func handleEmailLogin(c *gin.Context, email, password string) (string, error) {
	// Fetch the user from the database
	user, err := service.GetUserByEmail(c, email) // Assuming this is in your authService
	if err != nil {
		if err.Error() == "user not found" {
			return "", errors.New("invalid email or password")
		}
		return "", err // Database or other error
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate a JWT token
	return generateToken(user.Email, user.ID)
}

// handleGoogleLogin validates the Google login token and generates a JWT token
func handleGoogleLogin(c *gin.Context, token string) (string, error) {
	// Implement Google token verification here (use a library or API)
	// For example, using the Google API to verify the token

	// Mocked Google user data (replace with actual verification)
	googleUser := map[string]string{
		"email": "user@example.com", // Replace with email from the verified Google token
		"name":  "John Doe",
	}

	// Check if the Google user exists in the database
	user, err := service.GetUserByEmail(c, googleUser["email"])
	if err != nil {
		if err.Error() == "user not found" {
			// Register new user if not found
			user, err = service.CreateUser(c, googleUser["email"], googleUser["name"], "")
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	// Generate a JWT token
	return generateToken(user.Email, user.ID)
}

// generateToken generates a JWT token for the user
func generateToken(email, userID string) (string, error) {
	// Load the secret key from environment variables
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT secret key is not set")
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
