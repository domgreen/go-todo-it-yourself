package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func makeRoutes(todo Todo) http.Handler {
	ok := func(c *gin.Context) {
		c.String(http.StatusOK, "")
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

		c.Writer.Header().Set("Location", item.URL)
		c.JSON(http.StatusCreated, item)
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
		c.String(http.StatusNoContent, "")
	})

	routes.DELETE("/todos/:id", func(c *gin.Context) {
		todo.Delete(c.Params.ByName("id"))
		c.String(http.StatusNoContent, "")
	})

	return routes
}
