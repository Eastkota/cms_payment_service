package resolvers

import (
	"cms_payment_service/services"
	"cms_payment_service/model"
	
	"fmt"

	"github.com/graphql-go/graphql"
)

type CmsPaymentResolver struct {
	Services services.Services
}

func NewCmsPaymentResolver(service services.Services) *CmsPaymentResolver {

	return &CmsPaymentResolver{Services: service}
}

func (r *CmsPaymentResolver) GetCmsPaymentResponse(p graphql.ResolveParams) (interface{}, error) {
	offset, _ := p.Args["offset"].(int)
	limit, _ := p.Args["limit"].(int)

    paymentResponses, err := r.Services.GetRmaPaymentResponse(offset, limit)
    if err != nil {
        return nil, fmt.Errorf("error retrieving payment responses: %w", err)
    }
    return &model.GenericPaymentResponse{
		Data: &model.GenericRmaPaymentSuccessData{
			Message: paymentResponses,
		},
		Error: nil,
	}, nil
}

func (r *CmsPaymentResolver) GetStripePaymentResponse(p graphql.ResolveParams) (interface{}, error) {
	stripeResponses, err := r.Services.GetStripePaymentResponse()
	if err != nil {
		return nil, fmt.Errorf("error retrieving stripe payment responses: %w", err)
	}
	return &model.GenericPaymentResponse{
		Data: &model.GenericStripePaymentSuccessData{
			Message: stripeResponses,
		},
		Error: nil,
	}, nil
}

func (r *CmsPaymentResolver) GetInatePaymentResponse(p graphql.ResolveParams) (interface{}, error) {
	inateResponses, err := r.Services.GetInatePaymentResponse()
	if err != nil {
		return nil, fmt.Errorf("error retrieving inate payment responses: %w", err)
	}
	return &model.GenericPaymentResponse{
		Data: &model.GenericInatePaymentSuccessData{
			Message: inateResponses,
		},
		Error: nil,
	}, nil
}