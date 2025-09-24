package schema

import "github.com/graphql-go/graphql"

var GenericRmaPaymentSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericRmaPaymentSuccessResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GenericRmaPaymentSuccessData},
		"error": &graphql.Field{Type: RmaPaymentError},
	},
})
var GenericAeRmaPaymentSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAeRmaPaymentSuccessResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{Type: GenericAeRmaPaymentSuccessData},
		"error": &graphql.Field{Type: RmaPaymentError},
	},
})

var ArResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: ArResponseData},
		"error": &graphql.Field{Type: RmaPaymentError},
	},
})
