package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignUpController struct {
	SignUpUsecase domain.SignUpUsecase
	Env           *bootstrap.Env
}

func (sc *SignUpController) SignUp(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	w.Header().Set("Content-Type", "application/json")

	var signUpRequest domain.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&signUpRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignUpUsecase.GetUserByEmail(ctx, signUpRequest.Email)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "email already exists"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(r.FormValue("password")),
		bcrypt.DefaultCost,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "internal server error"})
		return
	}

	r.Form.Set("password", string(encryptedPassword))

	user := domain.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err = sc.SignUpUsecase.Create(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignUpUsecase.CreateRefreshToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := sc.SignUpUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupResponse)
}
