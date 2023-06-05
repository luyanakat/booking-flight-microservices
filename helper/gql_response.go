package helper

import "github.com/vektah/gqlparser/v2/gqlerror"

func GqlResponse(message string, code int) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    message,
		Extensions: map[string]interface{}{"statusCode": code},
	}
}
