package repository

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetById(ctx context.Context, id int64) (*entity.Transaction, error)
	GetByUser(ctx context.Context, id int64) ([]entity.Transaction, error)
	UpdateStatus(ctx context.Context, transactionID int64, newStatus string) error 
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *entity.Transaction) error {
	return r.db.WithContext(ctx).Create(&transaction).Error
}

func (r *transactionRepository) GetById(ctx context.Context, id int64) (*entity.Transaction, error){
	result := new(entity.Transaction)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (r *transactionRepository) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	result := make([]entity.Transaction, 0)
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *transactionRepository) GetByUser(ctx context.Context, id int64) ([]entity.Transaction, error){
	result := make([]entity.Transaction, 0)
	if err := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (r *transactionRepository) UpdateStatus(ctx context.Context, transactionID int64, newStatus string) error {
    // Memperbarui hanya kolom status berdasarkan id
    return r.db.WithContext(ctx).
        Model(&entity.Transaction{}).              
        Where("id = ?", transactionID).              
        Update("status", newStatus).                 
        Error
}