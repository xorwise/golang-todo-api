package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
	"github.com/xorwise/golang-todo-api/internal/utils"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewSignUpUsecase(ur domain.UserRepository, timeout time.Duration) domain.SignUpUsecase {
	return &signupUsecase{
		userRepository: ur,
		timeout:        timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.timeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.timeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
