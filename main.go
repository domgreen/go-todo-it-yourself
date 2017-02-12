package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	ID        string
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
	URL       string `json:"url"`
}

type Todo map[string]*TodoItem

func (t Todo) Create(item TodoItem, baseURL string) *TodoItem {
	item.ID = strconv.Itoa(1)
	item.URL = "http://" + baseURL + "/" + item.ID
	t[item.ID] = &item
	return &item
}

func (t Todo) GetAll() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

func (t Todo) Get(ID string) *TodoItem {
	return t[ID]
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
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	routes := gin.Default()
	routes.Use(cors)

	routes.OPTIONS("/todos", ok)

	routes.OPTIONS("/todos/:id", ok)

	routes.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todo.GetAll())
	})

	routes.GET("/todos/:id", func(c *gin.Context) {
		ID := c.Params.ByName("id")
		item := todo.Get(ID)
		c.JSON(http.StatusOK, item)
	})

	routes.POST("/todos", func(c *gin.Context) {
		template := TodoItem{}
		err := c.BindJSON(&template)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		baseURL := c.Request.Host + c.Request.URL.String()
		item := todo.Create(template, baseURL)
		c.JSON(http.StatusOK, item)
	})

	routes.DELETE("/todos", func(c *gin.Context) {
		todo.DeleteAll()
		c.String(http.StatusOK, "")
	})

	routes.Run(":" + port)
}
