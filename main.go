package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
}

type Todo map[string]*TodoItem

func (t Todo) Create(item TodoItem) {
	t["1"] = &item
}

func (t Todo) GetAll() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

func (t Todo) DeleteAll() {
	for k := range t {
		delete(t, k)
	}
}

func main() {
	todo := Todo{}

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
		c.JSON(http.StatusOK, todo.GetAll())
	})

	routes.POST("/todos", func(c *gin.Context) {
		template := TodoItem{}
		err := c.BindJSON(&template)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		todo.Create(template)
		c.JSON(http.StatusOK, template)
	})

	routes.DELETE("/todos", func(c *gin.Context) {
		todo.DeleteAll()
		c.String(http.StatusOK, "")
	})

	routes.Run(":" + port)
}
