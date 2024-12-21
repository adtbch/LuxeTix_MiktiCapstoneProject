package service

import (
	"context"
	"errors"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
)

type NotificationService interface {
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification dto.NotificationInput) error
	UserGetNotification(ctx context.Context, req dto.NotificationInput) ([]*entity.Notification, error)
}

type notificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(notificationRepository repository.NotificationRepository) NotificationService {
	return &notificationService{notificationRepository}
}

// Get All Notification ketika di get maka status notifikasi akan berubah menjadi true
func (s *notificationService) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	return s.notificationRepository.GetAllNotification(ctx)
}

// func untuk create notification
func (s *notificationService) CreateNotification(ctx context.Context, req dto.NotificationInput) error {
	userID := req.UserID
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	notification := &entity.Notification{
		UserID:  userID,
		Message: req.Message,
		IsRead:  false,
	}
	return s.notificationRepository.CreateNotification(ctx, notification)
}


// get notification after get chage value isRead to true and only get notification if isread false UserGetNotification
func (s *notificationService) UserGetNotification(ctx context.Context, req dto.NotificationInput) ([]*entity.Notification, error) {
	return s.notificationRepository.UserGetNotification(ctx, req)
}
