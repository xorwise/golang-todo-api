package domain

import (
	"context"
)

type Profile struct {
	ID    uint
	Name  string
	Email string
}

type ProfileResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID uint) (*Profile, error)
}
