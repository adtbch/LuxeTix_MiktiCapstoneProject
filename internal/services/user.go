package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.JWTCustomClaims, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, req dto.UpdateUserRequest) error
	Delete(ctx context.Context, user *entity.User) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	RequestResetPassword(ctx context.Context, username string) error
}

type userService struct {
	tokenService   TokenService
	cfg            *config.Config
	userRepository repository.UserRepository
}

func NewUserService(
	tokenService TokenService,
	cfg *config.Config,
	userRepository repository.UserRepository,
) UserService {
	return &userService{tokenService, cfg, userRepository}
}

func (s *userService) Login(ctx context.Context, email string, password string) (*entity.JWTCustomClaims, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("Email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Email atau password salah")
	}

	if user.IsVerified == false {
		return nil, errors.New("Silahkan verifikasi email terlebih dahulu")
	}

	expiredTime := time.Now().Add(time.Minute * 10)

	claims := &entity.JWTCustomClaims{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.Fullname,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Luxe Tix",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	return claims, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	user := new(entity.User)
	user.Email = req.Email
	user.Fullname = req.FullName
	user.Gender = req.Gender
	user.Username = req.Username
	user.Role = "User"
	user.Verify_token = utils.RandomString(16)
	user.IsVerified = false

	// Cek jika email sudah terdaftar
	exist, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err == nil && exist != nil {
		return errors.New("Email sudah digunakan")
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Simpan user ke database
	
	// Persiapkan email untuk verifikasi
	templatePath := "./templates/email/verify-email.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	
	var ReplacerEmail = struct {
		Token string
	}{
		Token: user.Verify_token,
	}
	
	var body bytes.Buffer
	if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
		return err
	}

	// Setup email message
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.SMTPConfig.Username)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Fast Tix : Verifikasi Email!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		s.cfg.SMTPConfig.Host,
		s.cfg.SMTPConfig.Port,
		s.cfg.SMTPConfig.Username,
		s.cfg.SMTPConfig.Password,
	)
	
	// Coba kirim email dan pastikan jika gagal, hapus user dari database
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	
	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *userService) GetById(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetById(ctx, id)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) error {
	user, err := s.userRepository.GetById(ctx, req.ID)
	if err != nil {
		return err
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	if req.FullName != "" {
		user.Fullname = req.FullName
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	return s.userRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, user *entity.User) error {
	return s.userRepository.Delete(ctx, user)
}

func (s *userService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	user, err := s.userRepository.GetByResetPasswordToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token reset password salah")
	}

	if req.Password == "" {
		return errors.New("Password tidak boleh kosong")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.Update(ctx, user)
}

func (s *userService) RequestResetPassword(ctx context.Context, email string) error {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("Email tersebut tidak ditemukan")
	}

	expiredTime := time.Now().Add(10 * time.Minute)

	claims := &entity.ResetPasswordClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			Issuer:    "Reset Password",
		},
	}

	token, err := s.tokenService.GenerateResetPasswordToken(ctx, *claims)
	if err != nil {
		return err
	}

	user.Reset_token = token
	err = s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	templatePath := "./templates/email/verify-email.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var ReplacerEmail = struct {
		Token string
	}{
		Token: user.Reset_token,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.SMTPConfig.Username)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Reset Password Request !")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		s.cfg.SMTPConfig.Host,
		s.cfg.SMTPConfig.Port,
		s.cfg.SMTPConfig.Username,
		s.cfg.SMTPConfig.Password,
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error {
	user, err := s.userRepository.GetByVerifyEmailToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token verifikasi email salah")
	}
	user.IsVerified = true
	return s.userRepository.Update(ctx, user)
}
