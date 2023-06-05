package middleware

import (
	"context"
	"mock-project/graphql/graph/model"
	"mock-project/helper"
	pb "mock-project/pb/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Auth(u pb.UserManagerClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.Next()
			return
		}

		tokenString := header
		// handlers validate in gRPC
		req := &pb.ParseTokenRequest{Token: tokenString}
		pRes, err := u.ParseToken(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Next()
			return
		}

		// put user in context
		userQ := &model.User{
			ID:         int(pRes.User.Id),
			Email:      pRes.User.Email,
			CustomerID: int(pRes.User.CustomerId),
			AccessID:   int(pRes.User.AccessId),
			CreatedAt:  helper.GetTimePointer(pRes.User.CreatedAt.AsTime()),
			UpdatedAt:  helper.GetTimePointer(pRes.User.UpdatedAt.AsTime()),
		}
		ctx := context.WithValue(c.Request.Context(), userCtxKey, userQ)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		return
	}
}

func GetUserFromContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
