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
	Error   string
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
		if err := sendOkResponse(w, Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetTodo - retrieves Todo by ID
func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}
	todo, err := h.Service.GetTodo(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Todo by ID", err)
		return
	}

	if err := sendOkResponse(w, todo); err != nil {
		panic(err)
	}

}

// GetAllTodos - retrieves all todos from the database
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		sendErrorResponse(w, "Failed to retrive all todos", err)
		return
	}

	if err := sendOkResponse(w, todos); err != nil {
		panic(err)
	}
}

// PostTodo - adds a new todo
func (h *Handler) PostTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	todo, err := h.Service.PostTodo(todo)

	if err != nil {
		sendErrorResponse(w, "Fail to post new todo", err)
		return
	}

	if err := sendOkResponse(w, todo); err != nil {
		panic(err)
	}
}

// UpdateTodo - Update a todo by ID
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
		return
	}

	todo, err = h.Service.UpdateTodo(uint(todoID), todo)

	if err != nil {
		sendErrorResponse(w, "Faill to update todo", err)
		return
	}

	if err := sendOkResponse(w, todo); err != nil {
		panic(err)
	}
}

// DeleteTodo - delete a todo by ID
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Faill to parse uint from ID", err)
		return
	}

	err = h.Service.DeleteTodo(uint(todoID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete todo by comment ID", err)
		return
	}

	if err = sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

// Error Helper Functions
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
