// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/juliendoutre/go-api-benchmark/data"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/juliendoutre/go-api-benchmark/swagger/models"
	"github.com/juliendoutre/go-api-benchmark/swagger/restapi/operations"
)

//go:generate swagger generate server --target ../../go-api-benchmark --name TodoList --spec ../specs.yml

var datastore = &data.Datastore{}

func formatTodo(todo *data.Todo) *models.Todo {
	return &models.Todo{
		Name: String(todo.Name),
		Tags: todo.Tags,
	}
}

func formatTodos(todos []*data.Todo) []*models.Todo {
	td := []*models.Todo{}

	for _, t := range todos {
		td = append(td, formatTodo(t))
	}

	return td
}

// String returns a pointer to a string
func String(str string) *string {
	return &str
}

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddTodoHandler = operations.AddTodoHandlerFunc(func(params operations.AddTodoParams) middleware.Responder {
		todo := params.Body

		err := datastore.AddTodo(*todo.Name, todo.Tags...)
		if err != nil {
			return operations.NewAddTodoBadRequest().WithPayload(&operations.AddTodoBadRequestBody{Message: String(err.Error())})
		}

		return operations.NewAddTodoOK().WithPayload(&operations.AddTodoOKBody{Message: String("ok")})
	})

	api.GetTodoByNameHandler = operations.GetTodoByNameHandlerFunc(func(params operations.GetTodoByNameParams) middleware.Responder {
		name := params.Name

		todo, err := datastore.GetTodoByName(name)
		if err != nil {
			return operations.NewGetTodoByNameNotFound().WithPayload(&operations.GetTodoByNameNotFoundBody{Message: String(err.Error())})
		}

		return operations.NewGetTodoByNameOK().WithPayload(&operations.GetTodoByNameOKBody{Todo: formatTodo(todo)})
	})

	api.GetTodosHandler = operations.GetTodosHandlerFunc(func(params operations.GetTodosParams) middleware.Responder {
		tagsString := *params.Tags

		var tags []string

		if tagsString == "" {
			tags = []string{}
		} else {
			tags = strings.Split(tagsString, ",")
		}

		todos := datastore.GetTodos(tags...)

		return operations.NewGetTodosOK().WithPayload(&operations.GetTodosOKBody{Todos: formatTodos(todos)})
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
