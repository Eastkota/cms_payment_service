package schema

import (
	"cms_payment_service/helpers"
	"cms_payment_service/model"
	"cms_payment_service/resolvers"

	"github.com/graphql-go/graphql"
)

// var query = (&queries.Query{Resolver: resolver})
func NewQueryType(resolver *resolvers.CmsPaymentResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
					if err != nil {
						return nil, err
					}

					serviceInfo := model.Service{
						Name:    "cmspaymentservice",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
			"GetRmaPaymentResponse": &graphql.Field{
				Type: GenericPaymentResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("list",resolver.GetCmsPaymentResponse))(p)
				},
			},
			"GetStripePaymentResponse": &graphql.Field{
				Type: GenericStripePaymentResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("list",resolver.GetStripePaymentResponse))(p)
				},
			},
		},
	})
}
