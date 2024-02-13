package domain

import (
	"context"
)

type ProfileResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID uint) (*User, error)
}
