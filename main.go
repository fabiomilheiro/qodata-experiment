// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Sample data
	todos = append(todos, Todo{ID: "1", Title: "Sample Todo", IsDone: false})

	// Route Handlers
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	log.Println("Server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
