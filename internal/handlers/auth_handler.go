package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sundayonah/phindcode_backend/internal/service"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

// SignInRequest defines the login request structure
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Method   string `json:"method"`
	Name     string `json:"name,omitempty"`
}

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// LogIn handles user login for both email/password and Google login
func (h *AuthHandler) LogIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	switch req.Method {
	case "email":
		token, name, err := h.handleEmailLogin(c, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "name": name})

	case "google":
		token, err := h.handleGoogleLogin(c, req.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login method"})
	}
}

// Register handles user registration (both email and Google)
func (h *AuthHandler) Register(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	switch req.Method {
	case "email":
		// If name is not provided, set a default or leave it empty
		if req.Name == "" {
			req.Name = "Default Name" // You can modify this to something else
		}

		token, err := h.handleEmailRegistration(c, req.Email, req.Password, req.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "name": req.Name})

	case "google":
		token, err := h.handleGoogleRegistration(c, req.Token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration method"})
	}
}

// handleEmailLogin validates the user's credentials and generates a JWT token
func (h *AuthHandler) handleEmailLogin(c *gin.Context, email, password string) (string, string, error) {
	user, err := h.svc.GetUserByEmail(c.Request.Context(), email) // Use AuthService
	if err != nil {
		if err.Error() == "user not found" {
			return "", "", errors.New("invalid email or password")
		}
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	userIDStr := strconv.Itoa(user.ID)

	token, err := generateToken(user.Email, userIDStr, user.IsAdmin)
	if err != nil {
		return "", "", err
	}

	return token, user.FullName, nil
}

// handleGoogleLogin validates the Google login token and generates a JWT token
func (h *AuthHandler) handleGoogleLogin(c *gin.Context, token string) (string, error) {
	// Verify the Google token using gin.Context
	payload, err := verifyGoogleToken(c, token)
	if err != nil {
		return "", errors.New("invalid Google token")
	}

	// Extract email and name from the token's claims
	email, ok := payload.Claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email in token payload")
	}

	name, ok := payload.Claims["name"].(string) // Extract name from the token
	if !ok {
		return "", errors.New("invalid name in token payload")
	}

	// Check if the Google user exists in the database
	user, err := h.svc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		if err.Error() == "user not found" {
			// Register new user if not found
			user, err = h.svc.CreateUser(c.Request.Context(), email, name, "") // Create with extracted name
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr, user.IsAdmin)
}

// handleEmailRegistration registers a new user with email and password and generates a JWT token
func (h *AuthHandler) handleEmailRegistration(c *gin.Context, email, password, name string) (string, error) {
	// Check if user already exists
	_, err := h.svc.GetUserByEmail(c.Request.Context(), email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create a new user in the database with the name
	user, err := h.svc.CreateUser(c.Request.Context(), email, name, string(hashedPassword))
	if err != nil {
		return "", err
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr, user.IsAdmin)
}

// handleGoogleRegistration registers a new user with Google login and generates a JWT token
func (h *AuthHandler) handleGoogleRegistration(c *gin.Context, token string) (string, error) {
	// Verify the Google token using gin.Context
	payload, err := verifyGoogleToken(c, token)
	if err != nil {
		return "", errors.New("invalid Google token")
	}

	// Extract email and name from the token's claims
	email, ok := payload.Claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email in token payload")
	}

	name, ok := payload.Claims["name"].(string)
	if !ok {
		return "", errors.New("invalid name in token payload")
	}

	// Check if the Google user exists in the database
	user, err := h.svc.GetUserByEmail(c.Request.Context(), email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// Register the new user with Google details
	user, err = h.svc.CreateUser(c.Request.Context(), email, name, "")
	if err != nil {
		return "", err
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr, user.IsAdmin)
}

func (h *AuthHandler) CreateAdmin(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user already exists
	_, err := h.svc.GetUserByEmail(c.Request.Context(), req.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create admin user using the service
	user, err := h.svc.CreateAdminUser(
		c.Request.Context(),
		req.Email,
		req.Name,
		string(hashedPassword),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin user"})
		return
	}

	userIDStr := strconv.Itoa(user.ID)

	// Generate JWT token
	token, err := generateToken(user.Email, userIDStr, user.IsAdmin)
	// token, err := generateToken(user.Email, strconv.Itoa(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.FullName,
		},
	})
}

func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	log.Println("GetAllUsers handler invoked") // Debug log
	users, err := h.svc.FetchAllUsers(c.Request.Context())
	if err != nil {
		log.Printf("Error fetching users: %v", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Convert to response format
	response := []gin.H{}
	for _, user := range users {
		response = append(response, gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.FullName,
		})
	}

	c.JSON(http.StatusOK, response)
}

func generateToken(email, userID string, isAdmin bool) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT secret key is not set")
	}
	claims := jwt.MapClaims{
		"email":   email,
		"userID":  userID,
		"isAdmin": isAdmin,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// generateToken generates a JWT token for the user
// func generateToken(email, userID string) (string, error) {
// secretKey := os.Getenv("JWT_SECRET")
// if secretKey == "" {
// 	return "", errors.New("JWT secret key is not set")
// }

// 	claims := jwt.MapClaims{
// 		"email":  email,
// 		"userID": userID,
// 		"exp":    time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secretKey))
// }

// verifyGoogleToken verifies the Google ID token
func verifyGoogleToken(c *gin.Context, token string) (*idtoken.Payload, error) {
	ctx := c.Request.Context()
	audience := os.Getenv("CLIENT_ID")
	if audience == "" {
		return nil, errors.New("CLIENT_ID environment variable is not set")
	}

	payload, err := idtoken.Validate(ctx, token, audience)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
