package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Gin router
	router := gin.Default()

	// Middleware: logger & recovery
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Static route
	router.GET("/halo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world",
		})
	})

	// Dynamic route dengan parameter path
	router.GET("/halo/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Halo, " + name + "!",
		})
	})

	// Login endpoint (POST + JSON body)
	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Bind body JSON ke struct
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}

		// Login logic
		if loginData.Email == "haikalfrastiawan16@gmail.com" && loginData.Password == "123" {
			c.JSON(200, gin.H{
				"message": "Login succeeded",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
	})

	// Query parameter endpoint
	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Query parameter 'name' is missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello, " + name,
		})
	})

	// Jalankan server di port 8000
	router.Run(":8000")
}
