package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	// Since the token is stateless, the logout functionality would typically involve
	// clearing the token on the client-side (browser, mobile).
	// No need to do anything on the server side for token-based auth.
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
