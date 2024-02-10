package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id uint) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
}
