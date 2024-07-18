package service

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context) ([]User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) Service {
	return &userService{db: db}
}

func (s *userService) CreateUser(ctx context.Context, user User) error {
	return s.db.Create(&user).Error
}

func (s *userService) GetUser(ctx context.Context, id string) (User, error) {
	var user User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user User) error {
	return s.db.Save(&user).Error
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.db.Delete(&User{}, "id = ?", id).Error
}

func (s *userService) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
