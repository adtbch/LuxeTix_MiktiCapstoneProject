package service

import (
	"context"
	"time"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	CreateTransaction(ctx context.Context, paymentRequest dto.PaymentRequest, req *entity.User) (string, error)
}

type paymentService struct {
	client              snap.Client
	cfg                 *config.Config
	notificationService NotificationService
}

func NewPaymentService(client snap.Client, cfg *config.Config, NotificationService NotificationService) PaymentService {
	return &paymentService{client, cfg, NotificationService}
}

func (s *paymentService) CreateTransaction(ctx context.Context, paymentRequest dto.PaymentRequest, req *entity.User) (string, error) {
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentRequest.OrderID,
			GrossAmt: paymentRequest.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Fullname,
			Email: req.Email,
		},
	}

	snapResponse, err := s.client.CreateTransaction(request)
	if err != nil {
		return "", err
	}

	// templatePath := "./templates/email/notif-submission.html"
	// tmpl, errs := template.ParseFiles(templatePath)
	// if errs != nil {
	// 	return "", err
	// }
	// var ReplacerEmail = struct {
	// 	Payment string
	// }{
	// 	Payment: snapResponse.RedirectURL,
	// }

	// var body bytes.Buffer
	// if err := tmpl.Execute(&body, ReplacerEmail); err != nil {
	// 	return "", err
	// }
	// m := gomail.NewMessage()
	// m.SetHeader("From", s.cfg.SMTPConfig.Username)
	// m.SetHeader("To", req.Email)
	// m.SetHeader("Subject", "LuxeTix : Payment Transaction ")
	// m.SetBody("text/html", body.String())

	// d := gomail.NewDialer(
	// 	s.cfg.SMTPConfig.Host,
	// 	s.cfg.SMTPConfig.Port,
	// 	s.cfg.SMTPConfig.Username,
	// 	s.cfg.SMTPConfig.Password,
	// )

	// if err := d.DialAndSend(m); err != nil {
	// 	return "",err
	// }

	notification := &dto.NotificationInput{
		UserID:    req.ID,
		Message:   "Your payment transaction has been submitted " + snapResponse.RedirectURL,
		Is_Read:   false,
		Create_at: time.Now(),
	}
	s.notificationService.CreateNotification(ctx, *notification)
	return snapResponse.RedirectURL, nil
}
