package domain

import "context"

type RefreshUsecase interface {
	GetUserByRefresh(c context.Context, refreshToken string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *User, secret string, expiry int) (string, error)
	UpdateUser(c context.Context, user *User) error
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
