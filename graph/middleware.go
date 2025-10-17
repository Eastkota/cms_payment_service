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

func PermissionMiddleware(actionName string, next ResolverFn) ResolverFn {
    // Both input 'next' and the return value must now be ResolverFn
    return func(p graphql.ResolveParams) (interface{}, error) {
        ctx := p.Context
        user := ctx.Value("user")
        
        // This check should ideally not fail if AuthMiddleware ran successfully,
        // but it's kept as a safety check.
        if user == nil {
            // Standard GraphQL error response (nil data, error object)
            return nil, fmt.Errorf("Unauthorized: User data missing from context")
        }

        userData, ok := user.(*model.User)
        if !ok {
             return nil, fmt.Errorf("Internal Error: Invalid user type in context")
        }
        
        hasPermission := false

        // --- Permission Check Logic (Unchanged) ---
        for _, role := range userData.Roles {
            if role.Name == "super admin" {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            for _, role := range userData.Roles {
                for _, permission := range role.Permissions {
                    if permission.Action.Action == actionName {
                        hasPermission = true
                        break 
                    }
                }
                if hasPermission {
                    break
                }
            }
        }
        // --- End Permission Check Logic ---

        if !hasPermission {
            fmt.Println("User does not have permission for action:", actionName)
            // Return a standard GraphQL error
            return nil, fmt.Errorf("Forbidden: Insufficient permissions for action '%s'", actionName)
        }

        // Permission granted, call the next resolver in the chain
        return next(p)
    }
}