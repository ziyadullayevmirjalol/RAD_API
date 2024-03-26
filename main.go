package main

import (
	"RADserver/handler"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// new router
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/users", handler.UsersHandler).Methods("GET","POST","PUT","DELETE")
	r.HandleFunc("/notes", handler.NotesHandler).Methods("GET","POST","PUT","DELETE")
	r.HandleFunc("/tasks", handler.TasksHandler).Methods("GET","POST","PUT","DELETE")

	// Define the allowed origins, methods, and headers
	allowedOrigins := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500", "http://localhost:5500","https://mirjalolziyadullayev.github.io"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	// CORS
	corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)

	fmt.Printf("Your server is running, Domain 'http://radserver.onrender.com', Port ':8000'\n")
	http.Handle("/", corsHandler)
	http.ListenAndServe(":8000", nil)
}
