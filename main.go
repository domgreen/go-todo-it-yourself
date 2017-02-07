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

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now())
	})

	r.Run(":" + port)
}
