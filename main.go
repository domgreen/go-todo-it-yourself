package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Title string `json:"title"`
	Order int    `json:"order"`
}

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

	routes := gin.Default()
	routes.Use(cors)

	routes.OPTIONS("/todos", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now().String())
	})

	routes.GET("/todos", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now().String())
	})

	routes.POST("/todos", func(c *gin.Context) {
		template := TodoItem{}
		err := c.BindJSON(&template)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, template)
	})

	routes.Run(":" + port)
}
