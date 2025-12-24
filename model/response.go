package model


type GenericRmaPaymentSuccessData struct {
	Message []PaymentResponse `json:"message"`
	Code    string `json:"code"`
}

type GenericStripePaymentSuccessData struct {
	Message []StripeResponse `json:"message"`
	Code    string            `json:"code"`
}

type GenericInatePaymentSuccessData struct {
	Message []InatePaymentResponse `json:"message"`
	Code    string                 `json:"code"`
}

type GenericPaymentResponse struct {
	Data  interface{}
	Error *CmsPaymentError
}
