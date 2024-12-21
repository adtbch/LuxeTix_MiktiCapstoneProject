package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
	"gopkg.in/gomail.v2"
)

type RequestEventService interface {
	GetAll(ctx context.Context, req dto.GetAllEventsSubmission) ([]entity.Event, error)
	GetByID(ctx context.Context, id int64) (*entity.Event, error)
	Approve(ctx context.Context, id int64) error
	Reject(ctx context.Context, id int64) error
	Cancel(ctx context.Context, submission *entity.Event, req dto.GetEventByIDRequest) error
	UpdateByUser(ctx context.Context, req dto.UpdateEventByUserRequest) error
	Create(ctx context.Context, req dto.CreateEventRequest) error
}

type requestEventService struct {
	cfg                    *config.Config
	requestEventRepository repository.RequestEventRepository
	transactionRepository  repository.TransactionRepository
	userRepository         repository.UserRepository
	notificationService    NotificationService
}

func NewSubmissionService(
	cfg *config.Config,
	requesteventRepository repository.RequestEventRepository,
	transactionRepository repository.TransactionRepository,
	userRepository repository.UserRepository,
	notification NotificationService,
) RequestEventService {
	return &requestEventService{
		cfg,
		requesteventRepository,
		transactionRepository,
		userRepository,
		notification,
	}
}

func (s *requestEventService) GetAll(ctx context.Context, req dto.GetAllEventsSubmission) ([]entity.Event, error) {
	return s.requestEventRepository.GetAll(ctx, req)
}

func (s *requestEventService) GetByID(ctx context.Context, id int64) (*entity.Event, error) {
	return s.requestEventRepository.GetByID(ctx, id)
}

func (s *requestEventService) UpdateByUser(ctx context.Context, req dto.UpdateEventByUserRequest) error {
	userID := req.UserID
	if userID == 0 {
		return errors.New("User ID tidak ditemukan")
	}
	submission, err := s.requestEventRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if submission.UserID != userID {
		return errors.New("Anda tidak memiliki hak untuk mengupdate pengajuan ini")
	}
	if submission.StatusRequest != "pending" {
		return errors.New("Pengajuan ini sudah tidak dapat diupdate karena sudah diterima atau ditolak oleh admin")
	}
	if req.Title != "" {
		submission.Title = req.Title
	}
	if req.Location != "" {
		submission.Location = req.Location
	}
	if req.Time != "" {
		submission.Time = req.Time
	}
	if req.Date != "" {
		submission.Date = req.Date
	}
	if req.Price != 0 {
		submission.Price = req.Price
	}
	if req.Description != "" {
		submission.Description = req.Description
	}
	if req.Category != "" {
		submission.Category = req.Category
	}
	if req.Quantity != 0 {
		submission.Quantity = req.Quantity
	}
	if req.Category != "" {
		submission.Category = req.Category
	}
	return s.requestEventRepository.Update(ctx, submission)
}

func (s *requestEventService) Approve(ctx context.Context, id int64) error {
	submission, err := s.requestEventRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if submission.StatusRequest != "pending" {
		return errors.New("Pengajuan ini sudah diterima atau ditolak oleh admin")
	}

	notification := &dto.NotificationInput{
		UserID:    submission.ID,
		Message:   "Your submission has been accepted! ",
		Is_Read:   false,
		Create_at: time.Now(),
	}
	s.notificationService.CreateNotification(ctx, *notification)

	submission.StatusRequest = "accepted"
	return s.requestEventRepository.Update(ctx, submission)
}

func (s *requestEventService) Reject(ctx context.Context, id int64) error {
	submission, err := s.requestEventRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if submission.StatusRequest != "pending" {
		return errors.New("Pengajuan ini sudah diterima atau ditolak oleh admin")
	}
	notification := &dto.NotificationInput{
		UserID:    submission.UserID,
		Message:   "Your submission has been rejected! ",
		Is_Read:   false,
		Create_at: time.Now(),
	}
	s.notificationService.CreateNotification(ctx, *notification)
	submission.StatusRequest = "rejected"
	return s.requestEventRepository.Update(ctx, submission)
}

func (s *requestEventService) Cancel(ctx context.Context, submission *entity.Event, req dto.GetEventByIDRequest) error {
	userID := req.UserID
	if userID == 0 {
		return errors.New("User ID tidak ditemukan")
	}
	submission, err := s.requestEventRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if submission.UserID != userID {
		return errors.New("Anda tidak memiliki hak untuk mengupdate pengajuan ini")
	}
	if submission.StatusRequest != "pending" {
		return errors.New("Pengajuan ini sudah tidak dapat dicancel karena sudah diterima atau ditolak oleh admin")
	}
	notification := &dto.NotificationInput{
		UserID:    userID,
		Message:   "Your submission has been canceled! ",
		Is_Read:   false,
		Create_at: time.Now(),
	}
	s.notificationService.CreateNotification(ctx, *notification)
	return s.requestEventRepository.Delete(ctx, submission)
}

func (s *requestEventService) Create(ctx context.Context, req dto.CreateEventRequest) error {
	// Ensure valid userID
	userID := req.UserID
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	user, err := s.userRepository.GetById(ctx, userID)
	if err != nil {
		return err
	}
	// Determine request status (paid/unpaid) based on price
	statusRequest := "unpaid"
	if req.Price == 0 {
		statusRequest = "pending"
		templatePath := "./templates/email/notif-submission.html"
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			return err
		}
		var body bytes.Buffer
		if err := tmpl.Execute(&body, nil); err != nil {
			return err
		}

		m := gomail.NewMessage()
		m.SetHeader("From", s.cfg.SMTPConfig.Username)
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", "LuxeTix: Submission Event Success!")
		m.SetBody("text/html", body.String())

		d := gomail.NewDialer(
			s.cfg.SMTPConfig.Host,
			s.cfg.SMTPConfig.Port,
			s.cfg.SMTPConfig.Username,
			s.cfg.SMTPConfig.Password,
		)
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}

		notification := &dto.NotificationInput{
			UserID:    userID,
			Message:   "Your submission waiting for admin approval! ",
			Is_Read:   false,
			Create_at: time.Now(),
		}
		s.notificationService.CreateNotification(ctx, *notification)
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

	// If price > 0, create a transaction
	if req.Price > 0 {
		eventID := event.ID
		amount := int64(float64(req.Price) * 0.2) // Example: 20% of event price

		transaction := &entity.Transaction{
			EventID:  eventID,
			UserID:   userID,
			Amount:   amount,
			Status:   "unpaid",
			Quantity: 1,
		}
		notification := &dto.NotificationInput{
			UserID:    userID,
			Message:   "Your create submission succecs, waiting for admin approval! ",
			Is_Read:   false,
			Create_at: time.Now(),
		}
		if err := s.requestEventRepository.Create(ctx, event); err != nil {
			return err
		}
		if err := s.notificationService.CreateNotification(ctx, *notification); err != nil {
			return err
		}
		// Create transaction
		if err := s.transactionRepository.Create(ctx, transaction); err != nil {
			return err
		}
	}

	return nil
}
