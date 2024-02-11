package domain

import (
	"context"
	"mime/multipart"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Password string
	Avatar   *string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id uint) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	Update(c context.Context, user *User) error
}

type UpdateUserRequest struct {
	Name   *string
	Email  *string
	Avatar *string
}

type UpdateUserUsecase interface {
	Update(c context.Context, user *User, req *UpdateUserRequest) error
	UploadAvatar(c context.Context, id uint, data multipart.File, handler *multipart.FileHeader) (string, error)
	GetUserByID(c context.Context, id uint) (User, error)
}
