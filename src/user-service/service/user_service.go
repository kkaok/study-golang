package service

import (
	"errors"
	"sync"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type userService struct {
	mu    sync.Mutex
	users map[string]User
}

func NewUserService() UserService {
	return &userService{
		users: make(map[string]User),
	}
}

func (s *userService) CreateUser(user User) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
	return user.ID, nil
}

func (s *userService) GetUser(id string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *userService) GetUsers() ([]User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}

func (s *userService) UpdateUser(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[user.ID]; !ok {
		return ErrUserNotFound
	}
	s.users[user.ID] = user
	return nil
}

func (s *userService) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return ErrUserNotFound
	}
	delete(s.users, id)
	return nil
}
