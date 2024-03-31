package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type FetchTaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *FetchTaskController) Fetch(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	taskRequest := domain.FetchTaskRequest{
		UserID: userID,
		Offset: offset,
		Limit:  limit,
	}

	tasks, err := tc.TaskUsecase.Fetch(r.Context(), taskRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	var tasksResponse []domain.FetchTaskResponse
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, domain.FetchTaskResponse{
			Title:       task.Title,
			Description: task.Description,
			Deadline:    task.Deadline,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			UserID:      task.UserID,
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasksResponse)
}
