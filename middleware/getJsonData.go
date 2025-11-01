package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware to read and store JSON body
func JSONBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
			if c.GetHeader("Content-Type") == "application/json" {
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err == nil && len(bodyBytes) > 0 {
					// Store original JSON body in context
					var jsonData map[string]interface{}
					if err := json.Unmarshal(bodyBytes, &jsonData); err == nil {
						c.Set("jsonData", jsonData)
					}

					// Reset body so handlers can still use BindJSON if needed
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}
		}
		c.Next()
	}
}
