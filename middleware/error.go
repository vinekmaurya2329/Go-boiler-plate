package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorResponse is a standard structure for API errors
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// HandleError sends a JSON response with dynamic message & status code
func HandleError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
	})
}

// GlobalErrorHandler - recovers from panics and prevents app crashes
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recover from any panic
		defer func() {
			if err := recover(); err != nil {
				
				fmt.Println("Panic recovered:", err)
				debug.PrintStack()

				// Respond with JSON (instead of crashing)
				c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
					Success: false,
					Message: "Internal Server Error",
					Error:   fmt.Sprintf("%v", err),
				})
			}
		}()

		// Continue request chain
		c.Next()
	}
}