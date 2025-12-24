package repositories

import "cms_payment_service/model"

type Repository interface {
	GetRmaPaymentResponse(offset, limit int) ([]model.PaymentResponse, error)
	GetStripePaymentResponse() ([]model.StripeResponse, error)
	GetInatePaymentResponse() ([]model.InatePaymentResponse, error)
}
