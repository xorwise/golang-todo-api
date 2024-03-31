package domain

import (
	"context"
	"mime/multipart"
)

type UpdateUserRequest struct {
	Name   *string
	Email  *string
	Avatar *string
}

type UpdateUserUsecase interface {
	Update(c context.Context, user *User, req *UpdateUserRequest) (*User, error)
	UploadAvatar(c context.Context, id uint, data multipart.File, handler *multipart.FileHeader) (string, error)
	GetUserByID(c context.Context, id uint) (User, error)
}

type UpdateUserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}
