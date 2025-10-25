package server

import (
	"awesomeProject1/internal/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.PutUser).Methods("PUT")

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("POST")
	http.ListenAndServe(":8080", r)
}
