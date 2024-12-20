package service

import (
	"context"
	"errors"
	"luxe/config"
	"luxe/internal/entity"
	"luxe/internal/http/dto"
	"luxe/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	Login(ctx context.Context, username, password string) (*entity.JWTCustomClaims, error)
}

type userService struct {
	cfg *config.Config
	userRepository repository.UserRepository
}

func NewUserService(cfg *config.Config, userRepository repository.UserRepository) UserService {
	return &userService{cfg, userRepository}
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	user := new(entity.User)
	user.Username = req.Username
	user.FullName = req.FullName
	user.Role = "User"
	user.Gender = req.Gender
	user.Email = req.Email
	
	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username tidak tersedia")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(ctx context.Context, username, password string) (*entity.JWTCustomClaims, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username atau password salah")
	}

	if user.IsVerified {
		return nil, errors.New("username atau password salah")
	}

	expiredTime := time.Now().Local().Add(time.Minute * 50)

	claims := &entity.JWTCustomClaims{
		Username: user.Username,
		Fullname: user.FullName,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "luxe_tix",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	return claims, nil
}

func (s *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetByID(ctx, id)
}

func (s *userService) Create(ctx context.Context, req dto.CreateUserByRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := &entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Role: req.Role,
	}
	return s.userRepository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) error {
	user, err := s.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	exist
}