package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juliendoutre/go-api-benchmark/data"
)

type h map[string]interface{}

func newGorillaMuxRouter(datastore *data.Datastore) *mux.Router {
	router := mux.NewRouter()

	router.StrictSlash(true)

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
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
	}).Methods("GET")

	router.HandleFunc("/todos/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

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
	}).Methods("GET")

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
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
	}).Methods("POST")

	return router
}
