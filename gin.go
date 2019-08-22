package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juliendoutre/go-api-benchmark/data"
)

func newGinRouter(datastore *data.Datastore) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/todos", func(c *gin.Context) {
		tagsString := c.Query("tags")

		var tags []string

		if tagsString == "" {
			tags = []string{}
		} else {
			tags = strings.Split(tagsString, ",")
		}

		todos := datastore.GetTodos(tags...)

		c.SecureJSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})

	router.GET("/todos/:name", func(c *gin.Context) {
		name := c.Param("name")

		todo, err := datastore.GetTodoByName(name)
		if err != nil {
			c.SecureJSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		} else {
			c.SecureJSON(http.StatusOK, gin.H{
				"todo": todo,
			})
		}
	})

	router.POST("/todos", func(c *gin.Context) {
		todoBytes, err := c.GetRawData()
		if err != nil {
			c.SecureJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			var todo data.Todo

			err = json.Unmarshal(todoBytes, &todo)
			if err != nil {
				c.SecureJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			} else {
				err = datastore.AddTodo(todo.Name, todo.Tags...)
				if err != nil {
					c.SecureJSON(http.StatusInternalServerError, gin.H{
						"message": err.Error(),
					})
				} else {
					c.SecureJSON(http.StatusOK, gin.H{
						"message": "ok",
					})
				}
			}
		}
	})

	return router
}
