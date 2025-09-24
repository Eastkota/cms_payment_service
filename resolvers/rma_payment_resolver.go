package resolvers

import (
	"rma_payment_service/helpers"
	"rma_payment_service/model"
	"rma_payment_service/services"
	
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type RmaPaymentResolver struct {
	Services services.Services
}

func NewRmaPaymentResolver(service services.Services) *RmaPaymentResolver {

	return &RmaPaymentResolver{Services: service}
}

func (r *RmaPaymentResolver) StorePaymentResponse(p graphql.ResolveParams) *model.GenericRmaPaymentResponse {
	var paymentResponseInput model.RmaPaymentResponseInput
	inputData := p.Args["input"].(map[string]interface{})
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = json.Unmarshal(jsonData, &paymentResponseInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = r.Services.StorePaymentResponse(paymentResponseInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericRmaPaymentResponse{
		Data: &model.GenericRmaPaymentSuccessData{
			Message: "payment stored successfully",
		},
	}
}

func (r *RmaPaymentResolver) MakeArRequest(p graphql.ResolveParams) *model.GenericRmaPaymentResponse {
	fmt.Println("MakeArRequest resolver called", p.Args)
	amount := p.Args["amount"].(float64)
	remitterEmail:= p.Args["remitter_email"].(string)
	product := p.Args["product"].(string)
	fmt.Println("MakeArRequest called with params:", amount, product, remitterEmail)
	err := helpers.ValidateArRequest(amount, product, remitterEmail)
	if err != nil {
		return helpers.FormatError(err)
	}
	data, err := r.Services.MakeArRequest(amount, product, remitterEmail)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericRmaPaymentResponse{
		Data: data,
	}
}

func (r *RmaPaymentResolver) MakeAeRequest(p graphql.ResolveParams) *model.GenericRmaPaymentResponse {
	remitterAccNo := p.Args["remitter_acc_no"].(string)
	txnId := p.Args["txn_id"].(string)
	remitterBankId := p.Args["remitter_bank_id"].(string)

	err := helpers.ValidateAeRequest(remitterAccNo, remitterBankId, txnId)
	if err != nil {
		return helpers.FormatError(err)
	}
	result, err := r.Services.MakeAeRequest(remitterAccNo, remitterBankId, txnId)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericRmaPaymentResponse{
		Data: &model.GenericAeRmaPaymentSuccessData{
			TxnId: result,
		},
		Error: nil,
	}
}

func (r *RmaPaymentResolver) MakeDrRequest(p graphql.ResolveParams) *model.GenericRmaPaymentResponse {
	txnId := p.Args["txn_id"].(string)
	remitterOtp := p.Args["remitter_otp"].(string)
	userId := p.Args["user_id"].(uuid.UUID)
	membershipDurationId := p.Args["membership_duration_id"].(uuid.UUID)

	err := helpers.ValidateDrRequest(txnId, remitterOtp, userId, membershipDurationId)
	if err != nil {
		return helpers.FormatError(err)
	}
	result, err := r.Services.MakeDrRequest(txnId, remitterOtp, userId, membershipDurationId)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericRmaPaymentResponse{
		Data: &model.GenericRmaPaymentSuccessData{
			Message: result,
		},
		Error: nil,
	}
}
