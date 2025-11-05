package services

import (
	"cms_payment_service/model"
)

type Services interface {
	GetRmaPaymentResponse(offset, limit int) ([]model.PaymentResponse, error)
	GetStripePaymentResponse() ([]model.StripeResponse, error)
}