package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
	"github.com/xorwise/golang-todo-api/internal/utils"
)

type refreshUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewRefreshUsecase(ur domain.UserRepository, timeout time.Duration) domain.RefreshUsecase {
	return &refreshUsecase{
		userRepository: ur,
		timeout:        timeout,
	}
}

func (ru *refreshUsecase) GetUserByRefresh(c context.Context, refreshToken string) (domain.User, error) {
	user, err := ru.userRepository.GetByRefresh(c, refreshToken)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ru *refreshUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (ru *refreshUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (ru *refreshUsecase) UpdateUser(c context.Context, user *domain.User) error {
	return ru.userRepository.Update(c, user)
}
