package repository

import (
	"context"
	"strings"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"gorm.io/gorm"
)

// EventRepository defines the interface for event-related database operations.
type RequestEventRepository interface {
	GetAll(ctx context.Context, req dto.GetAllEventsSubmission) ([]entity.Event, error)
	GetByID(ctx context.Context, id int64) (*entity.Event, error)
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, product *entity.Event) error
	Create(ctx context.Context, event *entity.Event) error
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
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllPending retrieves all pending events.
func (r *requestEventRepository) GetAll(ctx context.Context, req dto.GetAllEventsSubmission) ([]entity.Event, error) {
	result := make([]entity.Event, 0)
	query := r.db.WithContext(ctx)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Where("product_status = ? AND LOWER(product_name) LIKE ?", "pending", "%"+search+"%").
			Or("product_status = ? AND LOWER(product_category) LIKE ?", "pending", "%"+search+"%").
			Or("product_status = ? AND LOWER(product_address) LIKE ?", "pending", "%"+search+"%").
			Or("product_status = ? AND LOWER(product_price) LIKE ?", "pending", "%"+search+"%").
			Or("product_status = ? AND LOWER(product_date) LIKE ?", "pending", "%"+search+"%").
			Or("product_status = ? AND LOWER(product_time) LIKE ?", "pending", "%"+search+"%")
	}
	if req.Sort != "" && req.Order != "" {
		query = query.Order(req.Sort + " " + req.Order)
	}
	if req.Page != 0 && req.Limit != 0 {
		query = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}
	if err := query.Find(&result).Error; err != nil {
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

func (r *requestEventRepository) Delete(ctx context.Context, product *entity.Event) error {
	return r.db.WithContext(ctx).Delete(product).Error
}

func (r *requestEventRepository) Create(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Create(&event).Error
}