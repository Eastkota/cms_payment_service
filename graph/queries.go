package schema

import (
	"rma_payment_service/helpers"
	"rma_payment_service/model"
	"rma_payment_service/resolvers"
	"rma_payment_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// var query = (&queries.Query{Resolver: resolver})
func NewQueryType(resolver *resolvers.RmaPaymentResolver) *graphql.Object {
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
						Name:    "RmaPaymentService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
			"makeArRequest": &graphql.Field{
				Type: ArResponse,
				Args: graphql.FieldConfigArgument{
					"amount": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"remitter_email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"product": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.MakeArRequest(p), nil
				},
			},
			"makeAeRequest": &graphql.Field{
				Type: GenericAeRmaPaymentSuccessResponse,
				Args: graphql.FieldConfigArgument{
					"remitter_acc_no": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"txn_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"remitter_bank_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					return resolver.MakeAeRequest(p), nil
				},
			},
			"makeDrRequest": &graphql.Field{
				Type: GenericRmaPaymentSuccessResponse,
				Args: graphql.FieldConfigArgument{
					"txn_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"remitter_otp": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
					"membership_duration_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					return resolver.MakeDrRequest(p), nil
				},
			},
		},
	})
}
