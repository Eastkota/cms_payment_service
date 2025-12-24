package schema

import (
	"cms_payment_service/graph/scalar"
	
	"github.com/graphql-go/graphql"
)

var RmaPaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "RmaPaymentResponse",
	Fields: graphql.Fields{
		"BfsBfsTxnId":    &graphql.Field{Type: graphql.String},
		"BfsDebitAuthNo": &graphql.Field{Type: graphql.String},
		"BfsRemitterName": &graphql.Field{Type: graphql.String},
		"BfsTxnCurrency":  &graphql.Field{Type: graphql.String},
		"BfsBfsTxnTime":   &graphql.Field{Type: graphql.String},
		"BfsBenfId":      &graphql.Field{Type: graphql.String},
		"BfsRemitterBankId": &graphql.Field{Type: graphql.String},
		"BfsOrderNo":        &graphql.Field{Type: graphql.String},
		"BfsDebitAuthCode":  &graphql.Field{Type: graphql.String},
		"BfsTxnAmount":      &graphql.Field{Type: graphql.String},
		"BfsBenfTxnTime":    &graphql.Field{Type: graphql.String},
		"BfsMsgType":        &graphql.Field{Type: graphql.String},
		"MembershipDurationId": &graphql.Field{Type: scalar.UUID},
		"UserId":            &graphql.Field{Type: scalar.UUID},
	},
})

var StripePaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "StripePaymentResponse",
	Fields: graphql.Fields{
		"Amount":               &graphql.Field{Type: graphql.String},
		"TransactionId":        &graphql.Field{Type: graphql.String},
		"UserId":               &graphql.Field{Type: scalar.UUID},
		"MembershipDurationId": &graphql.Field{Type: scalar.UUID},
		"Status":               &graphql.Field{Type: graphql.String},
		"Remarks":              &graphql.Field{Type: graphql.String},
		"CreatedAt":            &graphql.Field{Type: graphql.DateTime},
		"UpdatedAt":            &graphql.Field{Type: graphql.DateTime},
		"Product":              &graphql.Field{Type: graphql.String},
	},
})

var GenericPaymentData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericPaymentData",
	Fields: graphql.Fields{
		"code":    &graphql.Field{Type: graphql.String},
		"Message": &graphql.Field{Type: graphql.String},
	},
})

var InatePaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "InatePaymentResponse",
	Fields: graphql.Fields{
		"ID":                   &graphql.Field{Type: scalar.UUID},
		"TransactionId":        &graphql.Field{Type: graphql.String},
		"Product":              &graphql.Field{Type: graphql.String},
		"PurchaseDate":         &graphql.Field{Type: graphql.DateTime},
		"UserId":               &graphql.Field{Type: scalar.UUID},
		"MembershipDurationId": &graphql.Field{Type: scalar.UUID},
		"Amount":               &graphql.Field{Type: graphql.Float},
		"Status":               &graphql.Field{Type: graphql.String},
	},
})
