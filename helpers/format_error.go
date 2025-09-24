package helpers

import "rma_payment_service/model"

func FormatError(err error) *model.GenericRmaPaymentResponse {
	return &model.GenericRmaPaymentResponse{
		Data: nil,
		Error: &model.RmaPaymentError{
			Message: err.Error(),
		},
	}
}
