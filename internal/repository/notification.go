package repository

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	GetAllNotification(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification *entity.Notification) error
	UserGetNotification(ctx context.Context, req dto.NotificationInput) ([]*entity.Notification, error)
	MarkNotificationAsRead(ctx context.Context, notificationID int64) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

// get all notification
func (r *notificationRepository) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	result := r.db.WithContext(ctx).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return Notifications, nil
}

// create notification
func (r *notificationRepository) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	result := r.db.WithContext(ctx).Create(&Notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// get notification after get change value isRead to true and only get notification if is_read false UserGetNotification
func (r *notificationRepository) UserGetNotification(ctx context.Context, req dto.NotificationInput) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	
	// Retrieve notifications with is_read = false
	result := r.db.WithContext(ctx).Where("is_read = ?", false).Where("user_id = ?", req.UserID).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mark retrieved notifications as read
	for _, notification := range Notifications {
		// Assuming you have a method to update the is_read field
		err := r.MarkNotificationAsRead(ctx, notification.ID)
		if err != nil {
			return nil, err
		}
	}

	return Notifications, nil
}

func (r *notificationRepository) MarkNotificationAsRead(ctx context.Context, notificationID int64) error {
	result := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", notificationID).Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
