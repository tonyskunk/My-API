package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func postUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "user created",
		"user":   user,
	})
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s", c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/api/hello", getHello)
	r.POST("/api/echo", postEcho)
	r.GET("/api/users/:id", getUser)
	r.POST("/api/users", postUser)

	r.Run(":8080")
}

func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
	})
}
func postEcho(c *gin.Context) {
	var json map[string]interface{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, json)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"name":    "John Doe",
	})
}

/*func postUser(c *gin.Context) {
	var user map[string]interface{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "user created",
		"user":   user,
	})
}*/
