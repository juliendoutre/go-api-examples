# Go Routing Frameworks Benchmark

Repository to tests several golang http routing frameworks:

- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorilla/Mux](https://github.com/gorilla/mux)
- [Go-swagger](https://github.com/go-swagger/go-swagger)

## Routes

- `GET /todos`

  - parameters:
    - `tags` (comma-separated values) (example: `?tags=tag1,tag2:tag2_value`)
  - response:

    ```json
    [
      {
        "name": "todo1",
        "tags": ["tag1", "tag2:tag2_value", "tag3"]
      },
      {
        "name": "todo2",
        "tags": ["tag1", "tag2:tag2_value"]
      }
    ]
    ```

- `GET /todos/:name`

  - response:

    ```json
    {
      "name": "todo2",
      "tags": ["tag1", "tag2:tag2_value"]
    }
    ```

- `POST /todos`

  - body:

    ```json
    {
      "todo": {
        "name": "todo3",
        "tags": ["tag3"]
      }
    }
    ```

See OpenAPI specifications in [`swaggers/specs.yml`](https://github.com/julien-doutre/go-api-benchmark/blob/master/swagger/specs.yml)
