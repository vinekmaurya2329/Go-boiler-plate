package main

import (
	"net/http"

	"boilerplate/db"
	"boilerplate/middleware"

	"github.com/gin-gonic/gin"
)
func main (){
	db.ConnectDB()
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