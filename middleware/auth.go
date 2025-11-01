package middleware

import (
	
	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {  
		token := c.GetHeader("Authorization")
		if token == "" {
			HandleError(c,401,"Authorization token required")
			return
		}
		//  check token validation  
            
	}
}