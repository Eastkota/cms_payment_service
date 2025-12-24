package schema

import "github.com/graphql-go/graphql"

var GenericRmaPaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenericRmaPaymentSuccessData",
    Fields: graphql.Fields{
        "Message": &graphql.Field{Type: graphql.NewList(RmaPaymentResponse)},
    },
})

var GenericStripePaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericStripePaymentSuccessData",
	Fields: graphql.Fields{
		"Message": &graphql.Field{Type: graphql.NewList(StripePaymentResponse)},
	},
})


// Define the new top-level object type
var GenericPaymentResponse = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenericPaymentResponse",
    Fields: graphql.Fields{
        "data": &graphql.Field{
            Type: GenericRmaPaymentSuccessData,
        },
        "error": &graphql.Field{
            Type: CmsPaymentError,
        },
    },
})
var GenericStripePaymentResponse = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenericStripePaymentResponse",
    Fields: graphql.Fields{
        "data": &graphql.Field{
            Type: GenericStripePaymentSuccessData,
        },
        "error": &graphql.Field{
            Type: CmsPaymentError,
        },
    },
})

var GenericInatePaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenericInatePaymentSuccessData",
    Fields: graphql.Fields{
        "Message": &graphql.Field{Type: graphql.NewList(InatePaymentResponse)},
    },
})

var GenericInatePaymentResponse = graphql.NewObject(graphql.ObjectConfig{
    Name: "GenericInatePaymentResponse",
    Fields: graphql.Fields{
        "data": &graphql.Field{
            Type: GenericInatePaymentSuccessData,
        },
        "error": &graphql.Field{
            Type: CmsPaymentError,
        },
    },
})

