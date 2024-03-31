package repository

import (
	"context"

	"github.com/xorwise/golang-todo-api/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	return ur.db.WithContext(c).Create(user).Error
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := ur.db.WithContext(c).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetByID(c context.Context, id uint) (domain.User, error) {
	var user domain.User
	if err := ur.db.WithContext(c).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	if err := ur.db.WithContext(c).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) Update(c context.Context, user *domain.User) error {
	return ur.db.WithContext(c).Save(user).Error
}

func (ur *userRepository) GetByRefresh(c context.Context, refreshToken string) (domain.User, error) {
	var user domain.User
	if err := ur.db.WithContext(c).Where("refresh_token = ?", refreshToken).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
