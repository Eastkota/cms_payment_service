package repositories

import "cms_payment_service/model"

type Repository interface {
	GetRmaPaymentResponse() ([]model.PaymentResponse, error)
	GetStripePaymentResponse() ([]model.StripeResponse, error)
}
