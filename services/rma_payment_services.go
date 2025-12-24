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

func (s *CmsPaymentService) GetRmaPaymentResponse(offset, limit int) ([]model.PaymentResponse, error) {
	return s.Repository.GetRmaPaymentResponse(offset, limit)
}
func (s *CmsPaymentService) GetStripePaymentResponse() ([]model.StripeResponse, error) {
	return s.Repository.GetStripePaymentResponse()
}

func (s *CmsPaymentService) GetInatePaymentResponse() ([]model.InatePaymentResponse, error) {
	return s.Repository.GetInatePaymentResponse()
}