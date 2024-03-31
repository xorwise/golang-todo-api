package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type UpdateTaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *UpdateTaskController) Update(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.ParseUint(r.FormValue("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	task, err := tc.TaskUsecase.GetByID(r.Context(), uint(taskID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if task.UserID != userID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You don't have permission to update this task")
		return
	}
	var taskRequest domain.UpdateTaskRequest

	err = json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = tc.TaskUsecase.Update(r.Context(), &task, &taskRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
