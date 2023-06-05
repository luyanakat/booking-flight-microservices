package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"log"
	"mock-project/graphql/graph"
	"mock-project/graphql/graph/model"
	"mock-project/helper"
	"mock-project/middleware"
	pb "mock-project/pb/proto"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User is the resolver for the User field.
func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *userOpsResolver) CreateUser(ctx context.Context, obj *model.UserOps, input *model.CreateUserInput) (*string, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, helper.GqlResponse("guest cannot create user", http.StatusUnauthorized)
	}

	if !helper.CheckAdmin(user) {
		return nil, helper.GqlResponse("user cannot create user", http.StatusUnauthorized)
	}
	pReq := &pb.User{
		Email:    input.Email,
		Password: input.Password,
		AccessId: int64(input.AccessID),
	}

	if err := helper.CheckLoginInput(input.Email, input.Password); err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusBadRequest)
	}

	pRes, err := r.userClient.CreateUser(ctx, pReq)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	id := strconv.Itoa(int(pRes.Id))

	return &id, nil
}

// UpdateUserAccessID is the resolver for the updateUserAccessId field.
func (r *userOpsResolver) UpdateUserAccessID(ctx context.Context, obj *model.UserOps, input model.UpdateAccessIDInput) (*string, error) {
	user := middleware.GetUserFromContext(ctx)
	if user == nil {
		return nil, helper.GqlResponse("guest cannot update user", http.StatusUnauthorized)
	}
	if helper.CheckAdmin(user) == false {
		return nil, helper.GqlResponse("user cannot update user", http.StatusUnauthorized)
	}

	md := metadata.Pairs("user", strconv.Itoa(user.ID))
	ctx = metadata.NewOutgoingContext(ctx, md)
	// send user id through metadata
	updatedUser, err := r.userClient.UpdateUser(ctx, &pb.UpdateUserRequest{
		AccessId: int64(input.AccessID),
	})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}
	resId := strconv.Itoa(int(updatedUser.Id))
	return &resId, nil
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *userOpsResolver) UpdateUserPassword(ctx context.Context, obj *model.UserOps, input *model.UpdateUserPasswordInput) (*string, error) {
	user := middleware.GetUserFromContext(ctx)
	if user == nil {
		return nil, helper.GqlResponse("guest cannot update user", http.StatusUnauthorized)
	}
	if helper.CheckAdmin(user) == false {
		return nil, helper.GqlResponse("user cannot update user", http.StatusUnauthorized)
	}
	if !helper.CheckPasswordLength(input.OldPassword) {
		return nil, helper.GqlResponse("password must be at least 8 characters", http.StatusBadRequest)
	}
	if !helper.CheckPasswordLength(input.Password) {
		return nil, helper.GqlResponse("new password must be at least 8 characters", http.StatusBadRequest)
	}
	md := metadata.Pairs("user", strconv.Itoa(user.ID))
	ctx = metadata.NewOutgoingContext(ctx, md)
	// send user id through metadata
	updatedUser, err := r.userClient.UpdateUserPassword(ctx, &pb.UpdateUserPasswordRequest{Password: input.Password, OldPassword: input.OldPassword})
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	resId := strconv.Itoa(int(updatedUser.Id))
	return &resId, nil
}

// RegisterCustomer is the resolver for the registerCustomer field.
func (r *userOpsResolver) RegisterCustomer(ctx context.Context, obj *model.UserOps, input *model.RegisterCustomer) (*model.CustomRegisterResponse, error) {
	date, err := time.Parse(time.RFC3339, input.DateOfBirth)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	if err := helper.CheckLoginInput(input.Email, input.Password); err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusBadRequest)
	}

	pReq := &pb.RegisterRequest{
		Email:          input.Email,
		Password:       input.Password,
		CustomerName:   input.CustomerName,
		Address:        input.Address,
		PhoneNumber:    input.PhoneNumber,
		IdentifyNumber: input.IdentifyNumber,
		DateOfBirth:    timestamppb.New(date),
		MemberCode:     *input.MemberCode,
	}

	pRes, err := r.userClient.RegisterCustomer(ctx, pReq)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	u := &model.User{
		ID:         int(pRes.UserId),
		CustomerID: int(pRes.CustomerId),
	}

	resQ := model.CustomRegisterResponse{
		Data:  u,
		Token: &pRes.JwtToken,
	}

	return &resQ, nil
}

// Login is the resolver for the login field.
func (r *userOpsResolver) Login(ctx context.Context, obj *model.UserOps, input *model.LoginInput) (*model.CustomLoginResponse, error) {
	if err := helper.CheckLoginInput(input.Email, input.Password); err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusBadRequest)
	}

	req := &pb.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	}
	loginRes, err := r.userClient.Login(ctx, req)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	resQ := &model.CustomLoginResponse{
		UserID: helper.GetIntPointer(int(loginRes.Id)),
		Token:  &loginRes.Token,
	}
	return resQ, nil
}

// Users is the resolver for the users field.
func (r *userQueryResolver) Users(ctx context.Context, obj *model.UserQuery, page *int, limit *int) (*model.CustomUserResponse, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, helper.GqlResponse("guest cannot create user", http.StatusUnauthorized)
	}

	if !helper.CheckAdmin(user) {
		return nil, helper.GqlResponse("user cannot create user", http.StatusUnauthorized)
	}

	req := &pb.ListUserRequest{
		Page:  int64(*page),
		Limit: int64(*limit),
	}

	res, err := r.userClient.ListUser(ctx, req)
	if err != nil {
		return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
	}

	var userRes []*model.UserResponse
	for _, user := range res.UserList {
		userQ := &model.UserResponse{}
		userQ.ID = int(user.Id)
		userQ.Email = user.Email
		userQ.Password = user.Password
		log.Println(user.CustomerId)
		if user.CustomerId == 0 {
			userQ.Customer = &model.Customer{}
		} else {
			cusRes, err := r.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: user.CustomerId})
			if err != nil {
				return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
			}

			customer := &model.Customer{
				ID:             int(cusRes.Id),
				Name:           cusRes.Name,
				Email:          cusRes.Email,
				Address:        cusRes.Address,
				PhoneNumber:    cusRes.PhoneNumber,
				IdentifyNumber: cusRes.IdentifyNumber,
				DateOfBirth:    cusRes.DateOfBirth.AsTime().Format("2006-01-02"),
				MemberCode:     cusRes.MemberCode,
				CreatedAt:      helper.GetTimePointer(cusRes.CreatedAt.AsTime()),
				UpdatedAt:      helper.GetTimePointer(cusRes.UpdatedAt.AsTime()),
			}
			userQ.Customer = customer
		}

		access, err := r.userClient.GetAccessLevel(ctx, &pb.GetAccessLevelRequest{Id: user.AccessId})
		if err != nil {
			return nil, helper.GqlResponse(err.Error(), http.StatusInternalServerError)
		}

		accessQ := &model.Access{
			ID:        int(access.Id),
			Name:      access.Name,
			CreatedAt: helper.GetTimePointer(access.CreatedAt.AsTime()),
			UpdatedAt: helper.GetTimePointer(access.UpdatedAt.AsTime()),
		}
		userQ.AccessType = accessQ
		userQ.CreatedAt = helper.GetTimePointer(user.CreatedAt.AsTime())
		userQ.UpdatedAt = helper.GetTimePointer(user.UpdatedAt.AsTime())

		userRes = append(userRes, userQ)
	}

	resQ := model.CustomUserResponse{
		Data:  userRes,
		Page:  helper.GetIntPointer(int(res.Page)),
		Limit: limit,
		Total: helper.GetIntPointer(int(res.Total)),
	}
	return &resQ, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// UserOps returns graph.UserOpsResolver implementation.
func (r *Resolver) UserOps() graph.UserOpsResolver { return &userOpsResolver{r} }

// UserQuery returns graph.UserQueryResolver implementation.
func (r *Resolver) UserQuery() graph.UserQueryResolver { return &userQueryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }