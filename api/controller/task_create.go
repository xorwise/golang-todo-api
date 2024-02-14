package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type CreateTaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *CreateTaskController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var taskRequest domain.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	task := &domain.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Deadline:    taskRequest.Deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserID:      userID,
	}

	err = tc.TaskUsecase.Create(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}
	taskResponse := domain.CreateTaskResponse{
		Title:       task.Title,
		Description: task.Description,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		UserID:      task.UserID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskResponse)
}
