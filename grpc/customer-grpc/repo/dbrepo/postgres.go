package dbrepo

import (
	"context"
	"mock-project/grpc/customer-grpc/ent"
	"mock-project/grpc/customer-grpc/ent/customer"
	"mock-project/grpc/customer-grpc/request"
)

func (m postgresDBRepo) CreateCustomer(ctx context.Context, c *ent.Customer) (*ent.Customer, error) {
	return m.client.Customer.Create().
		SetName(c.Name).
		SetEmail(c.Email).
		SetAddress(c.Address).
		SetPhoneNumber(c.PhoneNumber).
		SetIdentifyNumber(c.IdentifyNumber).
		SetDateOfBirth(c.DateOfBirth).
		SetMemberCode(c.MemberCode).
		Save(ctx)
}

func (m postgresDBRepo) GetCustomer(ctx context.Context, id int64) (*ent.Customer, error) {
	return m.client.Customer.Query().Where(customer.ID(id)).Only(ctx)
}

func (m postgresDBRepo) GetCustomerByEmail(ctx context.Context, email string) (*ent.Customer, error) {
	return m.client.Customer.Query().Where(customer.Email(email)).Only(ctx)
}

func (m postgresDBRepo) UpdateCustomer(ctx context.Context, id int64, c *ent.Customer) (*ent.Customer, error) {
	return m.client.Customer.UpdateOneID(id).
		SetName(c.Name).
		SetEmail(c.Email).
		SetAddress(c.Address).
		SetPhoneNumber(c.PhoneNumber).
		SetIdentifyNumber(c.IdentifyNumber).
		SetDateOfBirth(c.DateOfBirth).
		SetMemberCode(c.MemberCode).
		Save(ctx)
}

func (m postgresDBRepo) ListCustomer(ctx context.Context, paging request.Paging) ([]*ent.Customer, request.Paging, error) {
	total, err := m.client.Customer.Query().Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	customers, err := m.client.Customer.Query().Offset(int((paging.Page - 1) * paging.Limit)).Limit(int(paging.Limit)).All(ctx)
	if err != nil {
		return nil, paging, err
	}
	return customers, paging, nil
}

func (m postgresDBRepo) DeleteCustomer(ctx context.Context, id int64) error {
	return m.client.Customer.DeleteOneID(id).Exec(ctx)
}
