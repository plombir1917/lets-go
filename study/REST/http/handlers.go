package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest/todo"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandler(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), 400)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), 400)
	}
	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), 409)
		} else {
		}
	}
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {}
