package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
	Env            *bootstrap.Env
}

func (pc *ProfileController) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := pc.ProfileUsecase.GetProfileByID(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	profileResponse := domain.ProfileResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: *user.Avatar,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profileResponse)
}
