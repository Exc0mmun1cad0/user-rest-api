package userservice

import (
	"errors"
	"fmt"
	"test-api-task/internal/entity"
	userrepo "test-api-task/internal/repository/user"
)

type userRepo interface {
	GetUser(userID string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(userID string, newUserInfo *entity.User) (*entity.User, error)
	DeleteUser(userID string) error
}

type Service struct {
	userRepo userRepo
}

func NewService(userRepo userRepo) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) GetUser(userID string) (*entity.User, error) {
	const op = "service.userservice.GetUser"

	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		if errors.Is(err, userrepo.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Service) CreateUser(user *entity.User) (*entity.User, error) {
	const op = "service.userservice.CreateUser"

	newUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		if errors.Is(err, userrepo.ErrEmailAlreadyExists) {
			return nil, fmt.Errorf("%s: %w", op, ErrEmailAlreadyExists)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUser, nil
}

func (s *Service) UpdateUser(userID string, newUserInfo *entity.User) (*entity.User, error) {
	const op = "service.userservice.UpdateUser"

	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		if errors.Is(err, userrepo.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if newUserInfo.FirstName != "" {
		user.FirstName = newUserInfo.FirstName
	}
	if newUserInfo.LastName != "" {
		user.LastName = newUserInfo.LastName
	}
	if newUserInfo.Email != "" {
		user.Email = newUserInfo.Email
	}
	if newUserInfo.Age != 0 {
		user.Age = newUserInfo.Age
	}

	newUser, err := s.userRepo.UpdateUser(userID, user)
	if err != nil {
		if errors.Is(err, userrepo.ErrEmailAlreadyExists) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUser, nil
}

func (s *Service) DeleteUser(userID string) error {
	const op = "service.userservice.DeleteUser"

	err := s.userRepo.DeleteUser(userID)
	if errors.Is(err, userrepo.ErrUserNotFound) {
		return fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	return err
}
