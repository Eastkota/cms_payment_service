package helpers

import "cms_payment_service/model"

func FormatError(err error) *model.GenericPaymentResponse {
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.CmsPaymentError{
			Message: err.Error(),
		},
	}
}
