package model


type GenericRmaPaymentSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericAeRmaPaymentSuccessData struct {
	TxnId string `json:"txn_id"`
}

type GenericRmaPaymentResponse struct {
	Data  interface{}
	Error *RmaPaymentError
}