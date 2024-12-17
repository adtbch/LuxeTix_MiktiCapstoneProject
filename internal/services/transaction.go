package service

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type TransactionService interface {
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
}

type transactionService struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository}
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
