package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type DeleteTaskController struct {
	TaskUsecase domain.TaskUsecase
	Env         *bootstrap.Env
}

func (tc *DeleteTaskController) Delete(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
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
		json.NewEncoder(w).Encode("You don't have permission to delete this task")
		return
	}

	err = tc.TaskUsecase.Delete(r.Context(), &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
