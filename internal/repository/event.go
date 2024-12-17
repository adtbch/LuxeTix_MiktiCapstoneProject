package repository

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, event *entity.Event) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) (*entity.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetAll(ctx context.Context) ([]entity.Event, error) {
	result := make([]entity.Event, 0)
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *eventRepository) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	result := new(entity.Event)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *eventRepository) GetByIDPending(ctx context.Context, id int64) (*entity.Event, error) {
	result := new(entity.Event)
	if err := r.db.WithContext(ctx).Where("id = ?", id).Where("status = ?", "pending").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *eventRepository) GetAllPending(ctx context.Context) (*entity.Event, error) {
	result := new(entity.Event)
	if err := r.db.WithContext(ctx).Where("status = ?", "pending").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}
func (r *eventRepository) Create(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Create(&event).Error
}

func (r *eventRepository) Update(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Updates(&event).Error
}

func (r *eventRepository) Delete(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Delete(&event).Error
}
