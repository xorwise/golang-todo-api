package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
	"github.com/xorwise/golang-todo-api/internal/utils"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewLoginUsecase(ur domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		timeout:        timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.timeout)
	defer cancel()

	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (lu *loginUsecase) UpdateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, lu.timeout)
	defer cancel()

	return lu.userRepository.Update(ctx, user)
}
