package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"boilerplate/middleware"
)
func main (){
 r:= gin.Default()
 r.Use(middleware.JSONBodyMiddleware())
// Add logger + recovery middleware
r.Use(gin.Logger())

// Use your custom global error handler
r.Use(middleware.GlobalErrorHandler())

 r.GET("/ping",func (c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
 })

 r.Run()
}