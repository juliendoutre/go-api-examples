package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/juliendoutre/go-api-benchmark/swagger/restapi"
	"github.com/juliendoutre/go-api-benchmark/swagger/restapi/operations"
)

func main() {
	// datastore := &data.Datastore{}

	// ginRouter := newGinRouter(datastore)
	// ginRouter.Run(":8000")

	// gorillaRouter := newGorillaMuxRouter(datastore)
	// log.Fatal(http.ListenAndServe(":8001", gorillaRouter))

	// httpRouterRouter := newHTTPRouterRouter(datastore)
	// log.Fatal(http.ListenAndServe(":8002", httpRouterRouter))

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTodoListAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8003
	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
