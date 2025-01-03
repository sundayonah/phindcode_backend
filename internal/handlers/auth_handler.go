// package handlers

// import (
// 	"errors"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/sundayonah/phindcode_backend/internal/service"
// 	"golang.org/x/crypto/bcrypt"
// 	"google.golang.org/api/idtoken"
// )

// // SignInRequest defines the login request structure
// type SignInRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password,omitempty"`
// 	Token    string `json:"token,omitempty"`
// 	Method   string `json:"method"`
// }

// type AuthHandler struct {
// 	svc service.AuthService
// }

// func NewAuthHandler(svc service.AuthService) *AuthHandler {
// 	return &AuthHandler{svc: svc}
// }

// // LogIn handles user login for both email/password and Google login
// func (h *AuthHandler) LogIn(c *gin.Context) {
// 	var req SignInRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	switch req.Method {
// 	case "email":
// 		token, err := h.handleEmailLogin(c, req.Email, req.Password)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"token": token})

// 	case "google":
// 		token, err := h.handleGoogleLogin(c, req.Token)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"token": token})

// 	default:
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login method"})
// 	}
// }

// // handleEmailLogin validates the user's credentials and generates a JWT token
// func (h *AuthHandler) handleEmailLogin(c *gin.Context, email, password string) (string, error) {
// 	user, err := h.svc.GetUserByEmail(c.Request.Context(), email) // Use AuthService
// 	if err != nil {
// 		if err.Error() == "user not found" {
// 			return "", errors.New("invalid email or password")
// 		}
// 		return "", err
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		return "", errors.New("invalid email or password")
// 	}

// 	userIDStr := strconv.Itoa(user.ID)

// 	return generateToken(user.Email, userIDStr)
// }

// // handleGoogleLogin validates the Google login token and generates a JWT token
// func (h *AuthHandler) handleGoogleLogin(c *gin.Context, token string) (string, error) {
// 	// Verify the Google token using gin.Context
// 	payload, err := verifyGoogleToken(c, token)
// 	if err != nil {
// 		return "", errors.New("invalid Google token")
// 	}

// 	// Extract email and name from the token's claims
// 	email, ok := payload.Claims["email"].(string)
// 	if !ok {
// 		return "", errors.New("invalid email in token payload")
// 	}

// 	name, ok := payload.Claims["name"].(string)
// 	if !ok {
// 		return "", errors.New("invalid name in token payload")
// 	}

// 	// Check if the Google user exists in the database
// 	user, err := h.svc.GetUserByEmail(c.Request.Context(), email)
// 	if err != nil {
// 		if err.Error() == "user not found" {
// 			// Register new user if not found
// 			user, err = h.svc.CreateUser(c.Request.Context(), email, name, "")
// 			if err != nil {
// 				return "", err
// 			}
// 		} else {
// 			return "", err
// 		}
// 	}

// 	userIDStr := strconv.Itoa(user.ID)

// 	return generateToken(user.Email, userIDStr)
// }

// // generateToken generates a JWT token for the user
// func generateToken(email, userID string) (string, error) {
// 	secretKey := os.Getenv("JWT_SECRET")
// 	if secretKey == "" {
// 		return "", errors.New("JWT secret key is not set")
// 	}

// 	claims := jwt.MapClaims{
// 		"email":  email,
// 		"userID": userID,
// 		"exp":    time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secretKey))
// }

// // verifyGoogleToken verifies the Google ID token
// func verifyGoogleToken(c *gin.Context, token string) (*idtoken.Payload, error) {
// 	ctx := c.Request.Context()
// 	audience := os.Getenv("CLIENT_ID")
// 	if audience == "" {
// 		return nil, errors.New("CLIENT_ID environment variable is not set")
// 	}

// 	payload, err := idtoken.Validate(ctx, token, audience)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return payload, nil
// }

package handlers

import (
	"errors"
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
		token, err := h.handleEmailLogin(c, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

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
		token, err := h.handleEmailRegistration(c, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})

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
func (h *AuthHandler) handleEmailLogin(c *gin.Context, email, password string) (string, error) {
	user, err := h.svc.GetUserByEmail(c.Request.Context(), email) // Use AuthService
	if err != nil {
		if err.Error() == "user not found" {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr)
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

	name, ok := payload.Claims["name"].(string)
	if !ok {
		return "", errors.New("invalid name in token payload")
	}

	// Check if the Google user exists in the database
	user, err := h.svc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		if err.Error() == "user not found" {
			// Register new user if not found
			user, err = h.svc.CreateUser(c.Request.Context(), email, name, "")
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr)
}

// handleEmailRegistration registers a new user with email and password and generates a JWT token
func (h *AuthHandler) handleEmailRegistration(c *gin.Context, email, password string) (string, error) {
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

	// Create a new user in the database
	user, err := h.svc.CreateUser(c.Request.Context(), email, "", string(hashedPassword))
	if err != nil {
		return "", err
	}

	userIDStr := strconv.Itoa(user.ID)

	return generateToken(user.Email, userIDStr)
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

	return generateToken(user.Email, userIDStr)
}

// generateToken generates a JWT token for the user
func generateToken(email, userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT secret key is not set")
	}

	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

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
