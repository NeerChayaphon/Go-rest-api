package todo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - the struct for our todo service
type Service struct {
	DB *gorm.DB
}

// Todo -
type Todo struct {
	gorm.Model
	Name        string
	Description string
	Is_complete *bool
	Created     time.Time
}

// TodoService - the interface for the todo service
type TodoService interface {
	GetTodo(ID uint) (Todo, error)
	PostTodo(Todo Todo) (Todo, error)
	UpdateTodo(ID uint, newTodo Todo) (Todo, error)
	DeleteTodo(ID uint) error
	GetAllTodos() ([]Todo, error)
}

// NewService - returns a new todo service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetAllTodo - retrieves all todos from the database
func (s *Service) GetAllTodos() ([]Todo, error) {
	var todo []Todo
	if result := s.DB.Find(&todo); result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

// GetTodo - retrieves Todo by their ID from the database
func (s *Service) GetTodo(ID uint) (Todo, error) {
	var todo Todo
	if result := s.DB.First(&todo, ID); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// PostTodo - adds a new todo to the database
func (s *Service) PostTodo(todo Todo) (Todo, error) {
	if result := s.DB.Save(&todo); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// UpdateTodo - updates a todo by ID with new info
func (s *Service) UpdateTodo(ID uint, newTodo Todo) (Todo, error) {
	todo, err := s.GetTodo(ID)
	if err != nil {
		return Todo{}, err
	}

	if result := s.DB.Model(&todo).Updates(newTodo); result.Error != nil {
		return Todo{}, result.Error
	}

	return todo, nil
}

// DeleteTodo - deletes a todo from the database by ID
func (s *Service) DeleteTodo(ID uint) error {
	if result := s.DB.Delete(&Todo{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetTodosBySlug - retrieves all todos by slug (path - /article/name/)
// func (s *Service) GetTodosBySlug(slug string) ([]Todo, error) {
// 	var todos []Todo
// 	if result := s.DB.Find(&todos).Where("slug = ?", slug); result.Error != nil {
// 		return []Todo{}, result.Error
// 	}
// 	return todos, nil
// }
