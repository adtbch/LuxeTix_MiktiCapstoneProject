package service

import (
	"context"
	"errors"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type EventService interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	CreateEventByUser(ctx context.Context, req dto.CreateEventRequest) error
	Delete(ctx context.Context, event *entity.Event) error
	UpdateEventbyUser(ctx context.Context, req dto.UpdateEventByUserRequest) error
	UpdateEventbyAdmin(ctx context.Context, req dto.UpdateEventByAdminRequest) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) ([]entity.Event, error)
	SortFromExpensivestToCheapest(ctx context.Context) ([]entity.Event, error)
	SortFromCheapestToExpensivest(ctx context.Context) ([]entity.Event, error)
	SortNewestToOldest(ctx context.Context) ([]entity.Event, error)
	FilteringByCategory(ctx context.Context, category string) ([]entity.Event, error)
	FilteringByLocation(ctx context.Context, location string) ([]entity.Event, error)
	FilteringByPrice(ctx context.Context, price int64) ([]entity.Event, error)
	FilteringByDate(ctx context.Context, date string) ([]entity.Event, error)
	FilterMaxMinPrice(ctx context.Context, minPrice int64, maxPrice int64) ([]entity.Event, error)
}

type eventService struct {
	EventRepository       repository.EventRepository
	TransactionRepository repository.TransactionRepository
}

func NewEventService(eventRepository repository.EventRepository, transactionRepository repository.TransactionRepository) EventService {
	return &eventService{eventRepository, transactionRepository}
}

// Get all events
func (s *eventService) GetAll(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.GetAll(ctx)
}

// Get event by ID
func (s *eventService) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetById(ctx, id)
}

// Create event by user
func (s *eventService) CreateEventByUser(ctx context.Context, req dto.CreateEventRequest) error {
	// Ensure valid userID
	userID := req.UserID
	if userID == 0 {
		return errors.New("invalid user ID")
	}

	// Determine request status (paid/unpaid) based on price
	statusRequest := "unpaid"
	if req.Price == 0 {
		statusRequest = "paid"
	}

	// Create the event
	event := &entity.Event{
		Title:         req.Title,
		Location:      req.Location,
		Time:          req.Time,
		Date:          req.Date,
		Price:         req.Price,
		Description:   req.Description,
		StatusRequest: statusRequest,
		StatusEvent:   "available",
		UserID:        userID,
		Category:      req.Category,
		Quantity:      req.Quantity,
	}

	// Save event to database
	if err := s.EventRepository.Create(ctx, event); err != nil {
		return err
	}

	// If price > 0, create a transaction
	if req.Price > 0 {
		eventID := event.ID
		total := int64(float64(req.Price) * 0.2) // Example: 20% of event price

		transaction := &entity.Transaction{
			EventID:  eventID,
			UserID:   userID,
			Total:    total,
			Status:   "unpaid",
			Quantity: 1,
		}

		// Create transaction
		if err := s.TransactionRepository.Create(ctx, transaction); err != nil {
			return err
		}
	}

	return nil
}

// Update event by user
func (s *eventService) UpdateEventbyUser(ctx context.Context, req dto.UpdateEventByUserRequest) error {
	event, err := s.EventRepository.GetById(ctx, req.ID)
	if err != nil {
		return err
	}

	// Update fields if provided
	if req.Title != "" {
		event.Title = req.Title
	}
	if req.Location != "" {
		event.Location = req.Location
	}
	if req.Time != "" {
		event.Time = req.Time
	}
	if req.Date != "" {
		event.Date = req.Date
	}
	if req.Price != 0 {
		event.Price = req.Price
	}
	if req.Description != "" {
		event.Description = req.Description
	}
	if req.StatusEvent != "" {
		event.StatusEvent = req.StatusEvent
	}
	if req.Category != "" {
		event.Category = req.Category
	}

	// Save updated event
	return s.EventRepository.Update(ctx, event)
}

// Update event by admin
func (s *eventService) UpdateEventbyAdmin(ctx context.Context, req dto.UpdateEventByAdminRequest) error {
	event, err := s.EventRepository.GetById(ctx, req.ID)
	if err != nil {
		return err
	}

	// Admin updates event fields
	if req.Title != "" {
		event.Title = req.Title
	}
	if req.Location != "" {
		event.Location = req.Location
	}
	if req.Time != "" {
		event.Time = req.Time
	}
	if req.Date != "" {
		event.Date = req.Date
	}
	if req.Price != 0 {
		event.Price = req.Price
	}
	if req.Description != "" {
		event.Description = req.Description
	}
	if req.Category != "" {
		event.Category = req.Category
	}

	// Save updated event
	return s.EventRepository.Update(ctx, event)
}

// Delete event
func (s *eventService) Delete(ctx context.Context, event *entity.Event) error {
	return s.EventRepository.Delete(ctx, event)
}

// Get event by ID (pending)
func (s *eventService) GetByIDPending(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetByIDPending(ctx, id)
}

// Get all pending events
func (s *eventService) GetAllPending(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.GetAllPending(ctx)
}

// Sort events from expensive to cheapest
func (s *eventService) SortFromExpensivestToCheapest(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.SortFromExpensivestToCheapest(ctx)
}

// Sort events from cheapest to expensive
func (s *eventService) SortFromCheapestToExpensivest(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.SortFromCheapestToExpensivest(ctx)
}

// Sort events from newest to oldest
func (s *eventService) SortNewestToOldest(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.SortNewestToOldest(ctx)
}

// Filter events by category
func (s *eventService) FilteringByCategory(ctx context.Context, category string) ([]entity.Event, error) {
	return s.EventRepository.FilteringByCategory(ctx, category)
}

// Filter events by location
func (s *eventService) FilteringByLocation(ctx context.Context, location string) ([]entity.Event, error) {
	return s.EventRepository.FilteringByLocation(ctx, location)
}

// Filter events by price
func (s *eventService) FilteringByPrice(ctx context.Context, price int64) ([]entity.Event, error) {
	return s.EventRepository.FilteringByPrice(ctx, price)
}

// Filter events by date
func (s *eventService) FilteringByDate(ctx context.Context, date string) ([]entity.Event, error) {
	return s.EventRepository.FilteringByDate(ctx, date)
}

// Filter events by min and max price
func (s *eventService) FilterMaxMinPrice(ctx context.Context, minPrice int64, maxPrice int64) ([]entity.Event, error) {
	return s.EventRepository.FilteringByMinMaxPrice(ctx, minPrice, maxPrice)
}
