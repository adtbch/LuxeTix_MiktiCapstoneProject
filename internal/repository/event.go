package repository

import (
	"context"
	"strings"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"gorm.io/gorm"
)

// EventRepository defines the interface for event-related database operations.
type EventRepository interface {
	GetAll(ctx context.Context, req dto.GetAllEventRequest) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, event *entity.Event) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) ([]entity.Event, error)
	GetAllEventByOwner(ctx context.Context, id int64) ([]entity.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetAll(ctx context.Context, req dto.GetAllEventRequest) ([]entity.Event, error) {
	result := make([]entity.Event, 0)
	query := r.db.WithContext(ctx)

	// Pencarian berdasarkan title (case-insensitive)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Where("LOWER(title) LIKE ?", "%"+search+"%")
	}

	// Filter kategori (opsional, misalnya untuk kategori yang diinginkan)
	if req.Filter != "" {
		filter := strings.ToLower(req.Filter)
		query = query.Where("LOWER(category) = ?", filter)
	}

	// Filter tambahan dari kode pertama
	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}
	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}
	if req.Location != "" {
		query = query.Where("location LIKE ?", "%"+req.Location+"%")
	}
	if req.Time != "" {
		query = query.Where("time = ?", req.Time)
	}
	if req.Date != "" {
		query = query.Where("date = ?", req.Date)
	}
	if req.StatusEvent != "" {
		query = query.Where("status_event = ?", req.StatusEvent)
	}
	if req.StatusRequest != "" {
		query = query.Where("status_request = ?", req.StatusRequest)
	}
	if req.Price > 0 {
		query = query.Where("price = ?", req.Price)
	}

	// Sorting
	if req.Sort != "" && req.Order != "" {
		query = query.Order(req.Sort + " " + req.Order)
	}

	// Pagination
	if req.Page > 0 && req.Limit > 0 {
		query = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	// Eksekusi query
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// GetById retrieves a specific event by its ID.
func (r *eventRepository) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	var result entity.Event
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByIDPending retrieves a pending event by its ID.
func (r *eventRepository) GetByIDPending(ctx context.Context, id int64) (*entity.Event, error) {
	var result entity.Event
	if err := r.db.WithContext(ctx).Where("id = ?", id).Where("status = ?", "pending").First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllPending retrieves all pending events.
func (r *eventRepository) GetAllPending(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("status = ?", "pending").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// Create inserts a new event into the database.
func (r *eventRepository) Create(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Create(&event).Error
}

// Update modifies an existing event in the database.
func (r *eventRepository) Update(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Updates(&event).Error
}

// Delete removes an event from the database.
func (r *eventRepository) Delete(ctx context.Context, event *entity.Event) error {
	return r.db.WithContext(ctx).Delete(&event).Error
}

func (r *eventRepository) GetAllEventByOwner(ctx context.Context, id int64) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("owner_id = ?", id).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}