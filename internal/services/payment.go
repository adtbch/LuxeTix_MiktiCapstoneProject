package service

import (
	"context"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	CreateTransaction(ctx context.Context, paymentRequest *entity.Payment) (string, error)
}

type paymentService struct {
	client snap.Client
}

func NewPaymentService(client snap.Client) *paymentService {
	return &paymentService{client: client}
}

func (s *paymentService) CreateTransaction(ctx context.Context, paymentRequest *entity.Payment, req *entity.User) (string, error) {
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

	return snapResponse.RedirectURL, nil
}
