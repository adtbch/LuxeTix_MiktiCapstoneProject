package repository

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"gorm.io/gorm"
)

// EventRepository defines the interface for event-related database operations.
type RequestEventRepository interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetByID(ctx context.Context, id int64) (*entity.Event, error)
	Update(ctx context.Context, event *entity.Event) error
}

type requestEventRepository struct {
	db *gorm.DB
}

// NewEventRepository creates and returns a new instance of eventRepository.
func NewRequestEventRepository(db *gorm.DB) RequestEventRepository {
	return &requestEventRepository{db}
}

func (r *requestEventRepository) GetByID(ctx context.Context, id int64) (*entity.Event, error) {
	var result entity.Event
	if err := r.db.WithContext(ctx).Where("id = ?", id).Where("status = ?", "pending").First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllPending retrieves all pending events.
func (r *requestEventRepository) GetAll(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("status = ?", "pending").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *requestEventRepository) Update(ctx context.Context, event *entity.Event) error {
	var result entity.Event
	if err := r.db.WithContext(ctx).Where("status = ?", event.StatusRequest).Updates(&result).Error; err != nil {
		return err
	}
	return nil
}
