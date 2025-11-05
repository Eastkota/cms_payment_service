package repositories

import (
	"fmt"
	"cms_payment_service/model"
	"gorm.io/gorm"
)

type CmsPaymentRepository struct{
	DB *gorm.DB
}

func NewCmsPaymentRepository(db *gorm.DB) *CmsPaymentRepository {
	return &CmsPaymentRepository{DB: db}
}
func (r *CmsPaymentRepository) GetRmaPaymentResponse(offset, limit int) ([]model.PaymentResponse, error) {

	var paymentResponses []model.PaymentResponse
	result := r.DB.
		Offset(offset).
		Limit(limit).
		Find(&paymentResponses)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve payment responses: %v", result.Error)
	}

	return paymentResponses, nil
}

func (r *CmsPaymentRepository) GetStripePaymentResponse() ([]model.StripeResponse, error) {

	var stripeResponses []model.StripeResponse
	result := r.DB.Find(&stripeResponses)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve stripe payment responses: %v", result.Error)
	}

	return stripeResponses, nil
}