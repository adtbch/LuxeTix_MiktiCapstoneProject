package repository

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"gorm.io/gorm"
)

// EventRepository defines the interface for event-related database operations.
type EventRepository interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, event *entity.Event) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) ([]entity.Event, error)
	SortFromExpensivestToCheapest(ctx context.Context) ([]entity.Event, error)
	SortFromCheapestToExpensivest(ctx context.Context) ([]entity.Event, error)
	SortNewestToOldest(ctx context.Context) ([]entity.Event, error)
	FilteringByCategory(ctx context.Context, category string) ([]entity.Event, error)
	FilteringByLocation(ctx context.Context, location string) ([]entity.Event, error)
	FilteringByPrice(ctx context.Context, price int64) ([]entity.Event, error)
	FilteringByTime(ctx context.Context, time string) ([]entity.Event, error)
	FilteringByDate(ctx context.Context, date string) ([]entity.Event, error)
	FilteringByMinMaxPrice(ctx context.Context, minPrice int64, maxPrice int64) ([]entity.Event, error)
	SearchEvent(ctx context.Context, keyword string) ([]entity.Event, error)
}

// eventRepository is the implementation of EventRepository interface.
type eventRepository struct {
	db *gorm.DB
}

// NewEventRepository creates and returns a new instance of eventRepository.
func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

// GetAll retrieves all events from the database.
func (r *eventRepository) GetAll(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
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

// SortFromExpensivestToCheapest retrieves events sorted from the most expensive to the cheapest.
func (r *eventRepository) SortFromExpensivestToCheapest(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Order("price DESC").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// SortFromCheapestToExpensivest retrieves events sorted from the cheapest to the most expensive.
func (r *eventRepository) SortFromCheapestToExpensivest(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Order("price ASC").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// SortNewestToOldest retrieves events sorted from newest to oldest.
func (r *eventRepository) SortNewestToOldest(ctx context.Context) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByCategory retrieves events by category.
func (r *eventRepository) FilteringByCategory(ctx context.Context, category string) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("category = ?", category).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByLocation retrieves events by location.
func (r *eventRepository) FilteringByLocation(ctx context.Context, location string) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("location %like ?", location).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByPrice retrieves events by exact price.
func (r *eventRepository) FilteringByPrice(ctx context.Context, price int64) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("price = ?", price).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByTime retrieves events by exact time.
func (r *eventRepository) FilteringByTime(ctx context.Context, time string) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("time = ?", time).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByDate retrieves events by date.
func (r *eventRepository) FilteringByDate(ctx context.Context, date string) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("date = ?", date).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FilteringByMinMaxPrice retrieves events with a price between minPrice and maxPrice.
func (r *eventRepository) FilteringByMinMaxPrice(ctx context.Context, minPrice int64, maxPrice int64) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *eventRepository) SearchEvent(ctx context.Context, keyword string) ([]entity.Event, error) {
	var result []entity.Event
	if err := r.db.WithContext(ctx).Where("name %like ?", keyword).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}