package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "Content-Type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	r := gin.Default()
	r.Use(cors)

	r.GET("/todos", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now().String())
	})

	r.Run(":" + port)
}
