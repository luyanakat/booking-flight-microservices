package dbrepo

import (
	"context"
	"mock-project/grpc/user-grpc/ent"
	"mock-project/grpc/user-grpc/ent/accesslevel"
	"mock-project/grpc/user-grpc/ent/user"
	"mock-project/grpc/user-grpc/request"
)

func (m postgresDBRepo) CreateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	return m.client.User.Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetCustomerID(u.CustomerID).
		SetAccessID(u.AccessID).
		Save(ctx)
}

func (m postgresDBRepo) GetUser(ctx context.Context, id int64) (*ent.User, error) {
	return m.client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
}

func (m postgresDBRepo) GetUserByCustomerId(ctx context.Context, id int64) (*ent.User, error) {
	return m.client.User.Query().Where(user.CustomerID(id)).Only(ctx)
}

func (m postgresDBRepo) GetAccessLevel(ctx context.Context, id int64) (*ent.AccessLevel, error) {
	return m.client.AccessLevel.Query().Where(accesslevel.ID(id)).Only(ctx)
}

func (m postgresDBRepo) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return m.client.User.Query().Where(user.Email(email)).Only(ctx)
}

func (m postgresDBRepo) ListUser(ctx context.Context, paging request.Paging) ([]*ent.User, request.Paging, error) {
	total, err := m.client.User.Query().Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	users, err := m.client.User.Query().Offset(int((paging.Page - 1) * paging.Limit)).Limit(int(paging.Limit)).All(ctx)
	if err != nil {
		return nil, paging, err
	}
	return users, paging, nil
}

func (m postgresDBRepo) DeleteUser(ctx context.Context, id int64) error {
	return m.client.User.DeleteOneID(id).Exec(ctx)
}

func (m postgresDBRepo) UpdateUser(ctx context.Context, id int64, u *ent.User) (*ent.User, error) {
	return m.client.User.UpdateOneID(id).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetAccessID(u.AccessID).
		SetCustomerID(u.CustomerID).
		Save(ctx)
}

func (m postgresDBRepo) UpdateUserPassword(ctx context.Context, id int64, hashPass string) (*ent.User, error) {
	return m.client.User.UpdateOneID(id).
		SetPassword(hashPass).
		Save(ctx)
}
