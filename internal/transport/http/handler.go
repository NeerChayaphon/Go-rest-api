package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NeerChayaphon/go-rest-api/internal/todo"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our Todos service
type Handler struct {
	Router  *mux.Router
	Service *todo.Service
}

// Response - object to store responce from API
type Response struct {
	Message string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *todo.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/todo", h.GetAllTodos).Methods("GET")
	h.Router.HandleFunc("/api/todo", h.PostTodo).Methods("POST")
	h.Router.HandleFunc("/api/todo/{id}", h.GetTodo).Methods("GET")
	h.Router.HandleFunc("/api/todo/{id}", h.UpdateTodo).Methods("PUT")
	h.Router.HandleFunc("/api/todo/{id}", h.DeleteTodo).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetTodo - retrieves Todo by ID
func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from ID")
	}
	todo, err := h.Service.GetTodo(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error Retrieving Todo by ID")
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}

}

// GetAllTodos - retrieves all todos from the database
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	todos, err := h.Service.GetAllTodos()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrive all todos")
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// PostTodo - adds a new todo
func (h *Handler) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	todo, err := h.Service.PostTodo(todo)

	if err != nil {
		fmt.Fprintf(w, "Fail to post new todo")
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// UpdateTodo - Update a todo by ID
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	vars := mux.Vars(r)
	id := vars["id"]
	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	todo, err = h.Service.UpdateTodo(uint(todoID), todo)

	if err != nil {
		fmt.Fprintf(w, "Faill to update todo")
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// DeleteTodo - delete a todo by ID
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Faill to parse uint from ID")
	}

	err = h.Service.DeleteTodo(uint(todoID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete todo by comment ID")
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully delete todo"}); err != nil {
		panic(err)
	}
}
