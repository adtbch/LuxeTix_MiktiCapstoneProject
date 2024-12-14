package service

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type EventService interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	Create(ctx context.Context, req dto.CreateEventRequest) error
	Update(ctx context.Context, req dto.UpdateEventRequest) error
	Delete(ctx context.Context, event *entity.Event) error
}

type eventService struct {
	EventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{eventRepository}
}

func (s eventService) GetAll(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.GetAll(ctx)
}

func (s eventService) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetById(ctx, id)
}

func (s eventService) Create(ctx context.Context, req dto.CreateEventRequest) error {
	event := &entity.Event{
		Title:       req.Title,
		Location:    req.Location,
		Time:        req.Time,
		Date:        req.Date,
		Price:       req.Price,
		Description: req.Description,
		Status:      "pending",
		UserID:      req.UserID,
		Category:    req.Category,
	}
	return s.EventRepository.Create(ctx, event)
}

func (s eventService) Update(ctx context.Context, req dto.UpdateEventRequest) error {
	event, err := s.EventRepository.GetById(ctx, req.ID)
	if err != nil {
		return err
	}
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
	if req.Status != "" {
		event.Status = req.Status
	}
	if req.Category != "" {
		event.Category = req.Category
	}
	return s.EventRepository.Update(ctx, event)
}

func (s eventService) Delete(ctx context.Context, event *entity.Event) error {
	return s.EventRepository.Delete(ctx, event)
}
