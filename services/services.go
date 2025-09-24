package services

import (
	"rma_payment_service/model"

	"github.com/google/uuid"
)

type Services interface {
	MakeArRequest(amount float64, product, remitterEmail string)(*model.ArResponseData, error)
	MakeAeRequest(remitterAccNo, remitterBankId, product string) (string, error)
	MakeDrRequest(txnId, remitterOtp string, userId, membershipDurationId uuid.UUID) (string, error)

	StorePaymentResponse(paymentInput model.RmaPaymentResponseInput) (error)
}