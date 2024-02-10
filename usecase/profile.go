package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewProfileUsecase(ur domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: ur,
		timeout:        timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID uint) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.timeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.Profile{ID: user.ID, Name: user.Name, Email: user.Email}, err
}
