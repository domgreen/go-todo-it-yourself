package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TodoItem basic structure
type TodoItem struct {
	ID        string
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
	URL       string `json:"url"`
}

// Todo holding all the state
type Todo map[string]*TodoItem

// NextID Generates the for next id for Todo Items
func (t Todo) NextID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

// Create a new TodoItem and add it to the list
func (t Todo) Create(item TodoItem, baseURL string) *TodoItem {
	item.ID = t.NextID()
	item.URL = "http://" + baseURL + "/" + item.ID
	t[item.ID] = &item
	return &item
}

// GetAll will return all TodoItems
func (t Todo) GetAll() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

// Get will return a single item based on the Id
func (t Todo) Get(ID string) *TodoItem {
	return t[ID]
}

// Update single item
func (t Todo) Update(ID string, update TodoItem) {
	item := t.Get(ID)
	if len(update.Title) > 0 {
		item.Title = update.Title
	}

	if update.Completed != item.Completed {
		item.Completed = update.Completed
	}

	t[ID] = item
}

// DeleteAll removes all Todos
func (t Todo) DeleteAll() {
	for k := range t {
		delete(t, k)
	}
}

// Delete a single Todo
func (t Todo) Delete(ID string) {
	delete(t, ID)
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
		if item == nil {
			c.String(http.StatusNotFound, "")
		}
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

	routes.PATCH("/todos/:id", func(c *gin.Context) {
		template := TodoItem{}
		err := c.BindJSON(&template)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		fmt.Print(template.Title)
		fmt.Println(template.Order)

		todo.Update(c.Params.ByName("id"), template)
		item := todo.Get(c.Params.ByName("id"))
		c.JSON(http.StatusOK, item)
	})

	routes.DELETE("/todos", func(c *gin.Context) {
		todo.DeleteAll()
		c.String(http.StatusOK, "")
	})

	routes.DELETE("/todos/:id", func(c *gin.Context) {
		todo.Delete(c.Params.ByName("id"))
		c.String(http.StatusNotFound, "")
	})

	routes.Run(":" + port)
}
