package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
	"gopkg.in/gomail.v2"
)

type TransactionService interface {
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (error, *entity.Transaction)
	GetUserTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error)
	Update(ctx context.Context, transaction int64, newStatus string) error
}

type transactionService struct {
	cfg                   *config.Config
	TransactionRepository repository.TransactionRepository
	EventRepository       repository.EventRepository
	UserRepository        repository.UserRepository
	paymentService        PaymentService
}

func NewTransactionService(
	cfg *config.Config,
	transactionRepository repository.TransactionRepository,
	eventRepository repository.EventRepository,
	userRepository repository.UserRepository,
	paymentService PaymentService,
) TransactionService {
	return &transactionService{cfg, transactionRepository, eventRepository, userRepository, paymentService}
}

func (s *transactionService) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	return s.TransactionRepository.GetAll(ctx)
}

func (s *transactionService) GetById(ctx context.Context, id int64) (*entity.Transaction, error) {
	return s.TransactionRepository.GetById(ctx, id)
}

func (s *transactionService) Create(ctx context.Context, transaction *entity.Transaction) error {
	return s.TransactionRepository.Create(ctx, transaction)
}

func (s *transactionService) CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (error, *entity.Transaction) {
	var userID = req.UserID
	if userID == 0 {
		return errors.New("invalid user ID"), nil
	}

	user, err := s.UserRepository.GetById(ctx, userID)
	if err != nil {
		return err, nil
	}

	exist, err := s.EventRepository.GetById(ctx, req.EventID)
	if err != nil || exist == nil {
		return errors.New("Event not found"), nil
	}
	req.Amount = exist.Price

	amount := req.Amount * req.Quantity

	var status string
	if amount == 0 {
		status = "paid"
		templatePath := "./templates/email/notif-submission.html"
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			return err, nil
		}
		var ReplacerEmail = struct {
			Title    string
			Location string
			Date     string
			Price    int64
		}{
			Title:    exist.Title,
			Location: exist.Location,
			Date:     exist.Date,
			Price:    exist.Price,
		}

		var body bytes.Buffer
		if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
			return err, nil
		}
		m := gomail.NewMessage()
		m.SetHeader("From", s.cfg.SMTPConfig.Username)
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", "LuxeTix : Transaction Success!")
		m.SetBody("text/html", body.String())

		d := gomail.NewDialer(
			s.cfg.SMTPConfig.Host,
			s.cfg.SMTPConfig.Port,
			s.cfg.SMTPConfig.Username,
			s.cfg.SMTPConfig.Password,
		)
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}else{
		status = "unpaid"
	}

	transaction := &entity.Transaction{
		UserID:   userID,
		EventID:  req.EventID,
		Quantity: req.Quantity,
		Amount:   amount,
		Type:     "tiket",
		Status:   status,
	}

	err = s.TransactionRepository.Create(ctx, transaction)
	if err != nil {
		return err, nil // Kembalikan error dan nil jika gagal create
	}

	return nil, transaction
}

func (s *transactionService) GetUserTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	return s.TransactionRepository.GetByUser(ctx, userID)
}

func (s *transactionService) Update(ctx context.Context, transaction int64, newStatus string) error {
	return s.TransactionRepository.UpdateStatus(ctx, transaction, newStatus)
}
