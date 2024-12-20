package service

import (
	"context"
	"errors"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type TransactionService interface {
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest) error
	GetUserTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error)
}

type transactionService struct {
	TransactionRepository repository.TransactionRepository
	EventRepository       repository.EventRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, eventRepository repository.EventRepository) TransactionService {
	return &transactionService{transactionRepository, eventRepository}
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

func (s *transactionService) CreateOrder(ctx context.Context, req dto.CreateOrderRequest) error {
	var userID = req.UserID
	if userID == 0 {
		return errors.New("invalid user ID")
	}

	exist, err := s.EventRepository.GetById(ctx, req.EventID)
	if err != nil || exist == nil {
		return errors.New("Event not found")
	}
	req.Amount = exist.Price 

	amount := req.Amount * req.Quantity

	transaction := &entity.Transaction{
		UserID:   userID,
		EventID:  req.EventID,
		Quantity: req.Quantity,
		Amount:   amount,
		Status:   "unpaid",
	}
	return s.TransactionRepository.Create(ctx, transaction)
}

func (s *transactionService) GetUserTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error){
	return s.TransactionRepository.GetByUser(ctx, userID)
}