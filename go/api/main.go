package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/:ver", Logging(Index, "index")) // "Welcome!"
	router.GET("/:ver/todos", Logging(TodoIndex, "todo-index")) // get all the data
	router.GET("/:ver/todos/id/:todoId", Logging(TodoShow, "todo-show")) // get a data that has the id
	router.GET("/:ver/todos/notyet", Logging(TodoShowNotyet, "todo-show-not-yet")) // get all the data that has not been completed yet
	// POST DUE FORMAT
	// "due": "2015-09-15T14:00:12-00:00"    ,or
	// "due": "2015-09-15T14:00:13Z"    .
	router.POST("/:ver/todos", Logging(TodoCreate, "todo-create")) // post a data
	router.PATCH("/:ver/todos/completed/:todoId", Logging(TodoCompleted, "todo-completed")) // patch "completed" in todos
	router.DELETE("/:ver/todos/:todoId", Logging(TodoDelete, "todo-delete"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// API EXAMPLE
// curl -s -X GET localhost:8080/v1 | jq
// curl -s -X GET localhost:8080/v1/todos | jq
// curl -s -X GET localhost:8080/v1/todosid/1 | jq
// curl -s -X GET localhost:8080/v1/todos/notyet | jq
// curl -s -X POST localhost:8080/v1/todos -d '{"name": "sample", "completed": false, "due": "2018-09-15T17:45:13Z"}' | jq
// curl -s -X PATCH localhost:8080/v1/todos/completed/3 | jq
// curl -s -X DELETE localhost:8080/v1/todos/3 | jq
