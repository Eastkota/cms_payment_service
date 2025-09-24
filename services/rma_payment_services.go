package services

import (
	"cms_payment_service/model"
	"cms_payment_service/repositories"


)

type CmsPaymentService struct {
	Repository repositories.Repository // Inject Repository
}

func NewCmsPaymentService(repository repositories.Repository) *CmsPaymentService {
	return &CmsPaymentService{Repository: repository}
}

func (s *CmsPaymentService) GetRmaPaymentResponse() ([]model.PaymentResponse, error) {
	return s.Repository.GetRmaPaymentResponse()
}
func (s *CmsPaymentService) GetStripePaymentResponse() ([]model.StripeResponse, error) {
	return s.Repository.GetStripePaymentResponse()
}