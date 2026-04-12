package service

import (
	"context"
	"errors"
	"user-service/internal/config"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
}

type userService struct {
	repo   repository.UserRepository
	config config.Config // we'll inject config
}

func NewUserService(repo repository.UserRepository, cfg config.Config) UserService {
	return &userService{repo: repo, config: cfg}
}

func (s *userService) Register(ctx context.Context, req model.RegisterRequest) (*model.User, error) {
	existing, _ := s.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    req.Email,
		Password: string(hashed),
		FullName: req.FullName,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user, s.config.JWT.Secret, s.config.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
