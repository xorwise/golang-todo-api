package usecase

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
	"github.com/xorwise/golang-todo-api/internal/utils"
)

type updateUserUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewUpdateUserUsecase(ur domain.UserRepository, timeout time.Duration) domain.UpdateUserUsecase {
	return &updateUserUsecase{
		userRepository: ur,
		timeout:        timeout,
	}
}

func (uu *updateUserUsecase) Update(c context.Context, user *domain.User, req *domain.UpdateUserRequest) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Avatar != nil {
		user.Avatar = req.Avatar
	}

	err := uu.userRepository.Update(ctx, user)
	return user, err
}

func (uu *updateUserUsecase) UploadAvatar(c context.Context, id uint, data multipart.File, handler *multipart.FileHeader) (string, error) {
	path := "media/" + handler.Filename
	if !utils.IsMediaDirExists() {
		if err := utils.CreateMediaDir(); err != nil {
			return "", err
		}
	}
	dstFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, data)
	if err != nil {
		return "", err
	}

	return handler.Filename, nil

}

func (uu *updateUserUsecase) GetUserByID(c context.Context, id uint) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()
	return uu.userRepository.GetByID(ctx, id)
}
