package middleware

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/sundayonah/phindcode_backend/internal/service"
)

func AdminMiddleware(authSvc service.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get userID from the context (set by AuthMiddleware)
        userIDInterface, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        userIDStr, ok := userIDInterface.(string)
        if !ok {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
            c.Abort()
            return
        }

        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
            c.Abort()
            return
        }

        // Check if user is admin
        isAdmin, err := authSvc.IsAdmin(c.Request.Context(), userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking admin status"})
            c.Abort()
            return
        }

        if !isAdmin {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}