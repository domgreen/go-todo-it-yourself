package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Title string `json:"title"`
	Order int    `json:"order"`
}

func main() {
	port := os.Getenv("PORT")

	ok := func(c *gin.Context) {
		c.String(http.StatusOK, "")
	}

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

	routes.OPTIONS("/todos", ok)

	routes.GET("/todos", func(c *gin.Context) {
		c.String(http.StatusOK, "[]")
	})

	routes.POST("/todos", func(c *gin.Context) {
		template := TodoItem{}
		err := c.BindJSON(&template)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, template)
	})

	routes.DELETE("/todos", ok)

	routes.Run(":" + port)
}
