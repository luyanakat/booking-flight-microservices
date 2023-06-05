package repository

import (
	"context"
	"mock-project/grpc/customer-grpc/ent"
	"mock-project/grpc/customer-grpc/request"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, c *ent.Customer) (*ent.Customer, error)
	GetCustomer(ctx context.Context, id int64) (*ent.Customer, error)
	UpdateCustomer(ctx context.Context, id int64, f *ent.Customer) (*ent.Customer, error)
	ListCustomer(ctx context.Context, paging request.Paging) ([]*ent.Customer, request.Paging, error)
	DeleteCustomer(ctx context.Context, id int64) error
	GetCustomerByEmail(ctx context.Context, email string) (*ent.Customer, error)
}
