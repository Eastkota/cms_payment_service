package schema

import "github.com/graphql-go/graphql"

var GenericRmaPaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericRmaPaymentSuccessData",
	Fields: graphql.Fields{
		"code":    &graphql.Field{Type: graphql.String},
		"Message": &graphql.Field{Type: graphql.String},
	},
})


var GenericRmaPaymentSuccessDataResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericRmaPaymentSuccessDataResult",
	Fields: graphql.Fields{
		"generic_success_response": &graphql.Field{Type: GenericRmaPaymentSuccessData},
	},
})

var GenericAeRmaPaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAeRmaPaymentSuccessData",
	Fields: graphql.Fields{
		"txn_id":    &graphql.Field{Type: graphql.String},
	},
})

var Bank = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bank",
	Fields: graphql.Fields{
		"code": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

var ArResponseData = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArResponseData",
	Fields: graphql.Fields{
		"txnId": &graphql.Field{Type: graphql.String},
		"status": &graphql.Field{Type: graphql.String},
		"banks": &graphql.Field{Type: graphql.NewList(Bank)},
	},
})
