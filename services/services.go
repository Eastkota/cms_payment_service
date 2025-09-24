package services

import (
	"cms_payment_service/model"
)

type Services interface {
	GetRmaPaymentResponse() ([]model.PaymentResponse, error)
	GetStripePaymentResponse() ([]model.StripeResponse, error)
}