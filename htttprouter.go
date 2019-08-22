package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/juliendoutre/go-api-benchmark/data"
)

func newHTTPRouterRouter(datastore *data.Datastore) *httprouter.Router {
	router := httprouter.New()

	router.GET("/todos", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tagsParams := r.URL.Query()["tags"]

		var tagsString string

		if len(tagsParams) == 1 {
			tagsString = tagsParams[0]
		}

		var tags []string

		if tagsString == "" {
			tags = []string{}
		} else {
			tags = strings.Split(tagsString, ",")
		}

		todos := datastore.GetTodos(tags...)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h{
			"todos": todos,
		})
	})

	router.GET("/todos/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		name := params.ByName("name")

		todo, err := datastore.GetTodoByName(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(h{
				"message": err.Error(),
			})
		} else {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(h{
				"todo": todo,
			})
		}
	})

	router.POST("/todos", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		todoBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(h{
				"message": err.Error(),
			})
		} else {
			var todo data.Todo

			err = json.Unmarshal(todoBytes, &todo)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(h{
					"message": err.Error(),
				})
			} else {
				err = datastore.AddTodo(todo.Name, todo.Tags...)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(h{
						"message": err.Error(),
					})
				} else {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(h{
						"message": "ok",
					})
				}
			}
		}
	})

	return router
}
