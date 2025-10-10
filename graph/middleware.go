package schema

import (
    "cms_payment_service/helpers"
    "cms_payment_service/model"

    "fmt"
    "context"
    "net/http"

    "github.com/graphql-go/graphql"
)
type ResolverFn func(p graphql.ResolveParams) (interface{}, error)

func AuthMiddleware(next ResolverFn) ResolverFn { 
    return func(p graphql.ResolveParams) (interface{}, error) {
        ctx := p.Context
        userInterface := ctx.Value("user")

        var user *model.User
        var authError error 

        if userInterface == nil {
            if req, ok := ctx.Value("http_request").(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := helpers.ValidateToken(authHeader)
                if err != nil {
                    authError = err 
                } else if u == nil {
                    authError = fmt.Errorf("invalid_token")
                } else {
                    ctx = context.WithValue(ctx, "user", u)
                    p.Context = ctx
                    user = u
                }
            } else {
                authError = fmt.Errorf("invalid_token")
            }
        } else {
            user, _ = userInterface.(*model.User)
        }

        if user == nil && authError == nil {
            authError = fmt.Errorf("invalid_token")
        }

        if authError != nil {
            errResponse := helpers.FormatError(authError)
            return errResponse, nil 
        }

        return next(p)
    }
}