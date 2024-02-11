package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xorwise/golang-todo-api/api/middleware"
	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
)

type UpdateUserController struct {
	UpdateUserUsecase domain.UpdateUserUsecase
	Env               *bootstrap.Env
}

func (uc *UpdateUserController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(32 << 20) // 32 Mb
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "file is too large"})
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}
	defer file.Close()

	id, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := uc.UpdateUserUsecase.GetUserByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "user not found"})
		return
	}

	fileString, err := uc.UpdateUserUsecase.UploadAvatar(r.Context(), id, file, handler)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

	name, email := r.FormValue("name"), r.FormValue("email")
	userRequest := domain.UpdateUserRequest{
		Email:  &email,
		Name:   &name,
		Avatar: &fileString,
	}

	err = uc.UpdateUserUsecase.Update(r.Context(), &user, &userRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: err.Error()})
		return
	}

}
