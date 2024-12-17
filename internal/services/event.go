package service

import (
	"context"
	"fmt"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type EventService interface {
	GetAll(ctx context.Context) ([]entity.Event, error)
	GetById(ctx context.Context, id int64) (*entity.Event, error)
	CreateEventByUser(ctx context.Context, req dto.CreateEventRequest, tran dto.CreateEventTransactionRequest) error
	Delete(ctx context.Context, event *entity.Event) error
	UpdateEventbyUser(ctx context.Context, req dto.UpdateEventByUserRequest) error
	UpdateEventbyAdmin(ctx context.Context, req dto.UpdateEventByAdminRequest) error
	GetByIDPending(ctx context.Context, id int64) (*entity.Event, error)
	GetAllPending(ctx context.Context) (*entity.Event, error)
}

type eventService struct {
	EventRepository       repository.EventRepository
	TransactionRepository repository.TransactionRepository
}

func NewEventService(eventRepository repository.EventRepository, transactionRepository repository.TransactionRepository) EventService {
	return &eventService{eventRepository, transactionRepository}
}

func (s eventService) GetAll(ctx context.Context) ([]entity.Event, error) {
	return s.EventRepository.GetAll(ctx)
}

func (s eventService) GetById(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetById(ctx, id)
}

func (s *eventService) CreateEventByUser(ctx context.Context, req dto.CreateEventRequest, tran dto.CreateEventTransactionRequest) error {
	// Pastikan UserID yang diterima dari request sudah valid
	userID := req.UserID
	if userID == 0 {
		return fmt.Errorf("invalid UserID")
	}

	// Membuat event baru berdasarkan data dari request body
	event := &entity.Event{
		Title:         req.Title,
		Location:      req.Location,
		Time:          req.Time,
		Date:          req.Date,
		Price:         req.Price,
		Description:   req.Description,
		StatusRequest: "unpaid",
		StatusEvent:   "available",
		UserID:        userID, // UserID sudah dikirim dari handler
		Category:      req.Category,
		Quantity:      req.Quantity,
	}

	// Simpan event ke database
	if err := s.EventRepository.Create(ctx, event); err != nil {
		return err // Menangani error saat menyimpan event
	}

	// Ambil EventID yang baru dibuat
	eventID := event.ID // event.ID adalah ID yang dihasilkan oleh DB setelah event disimpan

	// Hitung total transaksi (misalnya, 20% dari harga event)
	total := int64(float64(req.Price) * 0.2) // Pastikan perhitungan yang tepat

	// Membuat transaksi dengan EventID yang baru
	transaction := &entity.Transaction{
		EventID:  eventID,
		UserID:   userID,
		Total:    total,
		Status:   "unpaid",
		Quantity: 1,
	}

	// Simpan transaksi ke database
	if err := s.TransactionRepository.Create(ctx, transaction); err != nil {
		return err // Menangani error saat menyimpan transaksi
	}

	return nil
}

func (s eventService) UpdateEventbyUser(ctx context.Context, req dto.UpdateEventByUserRequest) error {
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
	if req.StatusEvent != "" {
		event.StatusEvent = req.StatusEvent
	}
	if req.Category != "" {
		event.Category = req.Category
	}
	return s.EventRepository.Update(ctx, event)
}

func (s eventService) UpdateEventbyAdmin(ctx context.Context, req dto.UpdateEventByAdminRequest) error {
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
	if req.Category != "" {
		event.Category = req.Category
	}
	return s.EventRepository.Update(ctx, event)
}

func (s eventService) Delete(ctx context.Context, event *entity.Event) error {
	return s.EventRepository.Delete(ctx, event)
}

func (s eventService) GetByIDPending(ctx context.Context, id int64) (*entity.Event, error) {
	return s.EventRepository.GetByIDPending(ctx, id)
}

func (s eventService) GetAllPending(ctx context.Context) (*entity.Event, error) {
	return s.EventRepository.GetAllPending(ctx)
}
