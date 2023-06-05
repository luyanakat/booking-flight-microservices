package repository

import (
	"context"
	"mock-project/grpc/user-grpc/ent"
	"mock-project/grpc/user-grpc/request"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
	GetUser(ctx context.Context, id int64) (*ent.User, error)
	ListUser(ctx context.Context, paging request.Paging) ([]*ent.User, request.Paging, error)
	GetAccessLevel(ctx context.Context, id int64) (*ent.AccessLevel, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, id int64, u *ent.User) (*ent.User, error)
	GetUserByCustomerId(ctx context.Context, id int64) (*ent.User, error)
	UpdateUserPassword(ctx context.Context, id int64, hashPass string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
}
