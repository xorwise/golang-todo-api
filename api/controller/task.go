package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	t, err := time.Parse("2006-01-02 15:04", r.FormValue("deadline"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}
	taskRequest := domain.CreateTaskRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Deadline:    t,
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
	w.WriteHeader(http.StatusCreated)
}
