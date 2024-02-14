package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type RefreshController struct {
	RefreshUsecase domain.RefreshUsecase
	Env            *bootstrap.Env
}

func (rc *RefreshController) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var refreshRequest domain.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&refreshRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := rc.RefreshUsecase.GetUserByRefresh(r.Context(), refreshRequest.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := rc.RefreshUsecase.CreateAccessToken(&user, rc.Env.AccessTokenSecret, rc.Env.AccessTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rc.RefreshUsecase.CreateRefreshToken(&user, rc.Env.RefreshTokenSecret, rc.Env.RefreshTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	user.RefreshToken = &refreshToken
	err = rc.RefreshUsecase.UpdateUser(r.Context(), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshResponse := domain.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refreshResponse)

}
