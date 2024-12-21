package service

import (
	"context"
	

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type EventService interface {
	GetAll(ctx context.Context, req dto.GetAllEventRequest) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	
	Delete(ctx context.Context, event *entity.Event) error
	UpdateEventbyUser(ctx context.Context, req dto.UpdateEventByUserRequest) error
	UpdateEventbyAdmin(ctx context.Context, req dto.UpdateEventByAdminRequest) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) ([]entity.Event, error)
	GetAllEventByOwner(ctx context.Context, id int64) ([]entity.Event, error)
}

type eventService struct {
	EventRepository       repository.EventRepository
	TransactionRepository repository.TransactionRepository
}

func NewEventService(eventRepository repository.EventRepository, transactionRepository repository.TransactionRepository) EventService {
	return &eventService{eventRepository, transactionRepository}
}

// Get all events
func (s *eventService) GetAll(ctx context.Context, req dto.GetAllEventRequest) ([]entity.Event, error) {
	return s.EventRepository.GetAll(ctx, req)
}

// Get event by ID
func (s *eventService) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetById(ctx, id)
}

// Create event by user

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

func (s *eventService) GetAllEventByOwner(ctx context.Context, id int64) ([]entity.Event, error) {
	return s.EventRepository.GetAllEventByOwner(ctx, id)
}

// Sort events from expensive to cheapest
