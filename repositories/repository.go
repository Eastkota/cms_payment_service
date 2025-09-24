package repositories

import "rma_payment_service/model"

type Repository interface {
	StorePaymentResponse(paymentInput model.RmaPaymentResponseInput) error
}
