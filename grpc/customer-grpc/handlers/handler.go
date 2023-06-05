package handlers

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mock-project/grpc/customer-grpc/ent"
	"mock-project/grpc/customer-grpc/internal/auth"
	repository "mock-project/grpc/customer-grpc/repo"
	"mock-project/grpc/customer-grpc/request"
	pb "mock-project/pb/proto"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerManagerServer
	customerRepo repository.CustomerRepository
	userClient   pb.UserManagerClient
}

func NewCustomerHandler(customerRepo repository.CustomerRepository, userClient pb.UserManagerClient) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepo: customerRepo,
		userClient:   userClient,
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer := &ent.Customer{
		Name:           req.Name,
		Email:          req.Email,
		Address:        req.Address,
		PhoneNumber:    req.PhoneNumber,
		IdentifyNumber: req.IdentifyNumber,
		DateOfBirth:    req.DateOfBirth.AsTime(),
		MemberCode:     req.MemberCode,
	}

	c, err := h.customerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Customer{
		Id:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		Address:        c.Address,
		PhoneNumber:    c.PhoneNumber,
		IdentifyNumber: c.IdentifyNumber,
		DateOfBirth:    timestamppb.New(c.DateOfBirth),
		MemberCode:     c.MemberCode,
	}
	return pRes, nil

}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer := &ent.Customer{
		Name:           req.Name,
		Email:          req.Email,
		Address:        req.Address,
		PhoneNumber:    req.PhoneNumber,
		IdentifyNumber: req.IdentifyNumber,
		DateOfBirth:    req.DateOfBirth.AsTime(),
		MemberCode:     req.MemberCode,
	}

	c, err := h.customerRepo.GetCustomerByEmail(ctx, customer.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, err
	}

	customerUpdate, err := h.customerRepo.UpdateCustomer(ctx, c.ID, customer)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Customer{
		Id:             customerUpdate.ID,
		Name:           customerUpdate.Name,
		Email:          customerUpdate.Email,
		Address:        customerUpdate.Address,
		PhoneNumber:    customerUpdate.PhoneNumber,
		IdentifyNumber: customerUpdate.IdentifyNumber,
		DateOfBirth:    timestamppb.New(customerUpdate.DateOfBirth),
		MemberCode:     customerUpdate.MemberCode,
	}

	return pRes, nil
}

func (h *CustomerHandler) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.Customer, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	c, err := h.customerRepo.GetCustomer(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, err
	}

	pRes := &pb.Customer{
		Id:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		Address:        c.Address,
		PhoneNumber:    c.PhoneNumber,
		IdentifyNumber: c.IdentifyNumber,
		DateOfBirth:    timestamppb.New(c.DateOfBirth),
		MemberCode:     c.MemberCode,
		CreatedAt:      timestamppb.New(c.CreatedAt),
		UpdatedAt:      timestamppb.New(c.UpdatedAt),
	}

	return pRes, nil
}

func (h *CustomerHandler) GetCustomerByEmail(ctx context.Context, req *pb.GetCustomerByEmailRequest) (*pb.Customer, error) {
	c, err := h.customerRepo.GetCustomerByEmail(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, err
	}

	pRes := &pb.Customer{}
	if err := copier.Copy(&pRes, c); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pRes.Id = c.ID

	return pRes, nil
}

func (h *CustomerHandler) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (*pb.ListCustomerResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}

	list, pg, err := h.customerRepo.ListCustomer(ctx, paging)
	if err != nil {
		return nil, err
	}

	// append pb customer to list response
	var customerPbRes []*pb.Customer
	for _, customer := range list {
		customerPb := &pb.Customer{}
		customerPb.Id = customer.ID
		customerPb.Name = customer.Name
		customerPb.Email = customer.Email
		customerPb.Address = customer.Address
		customerPb.PhoneNumber = customer.PhoneNumber
		customerPb.IdentifyNumber = customer.IdentifyNumber
		customerPb.DateOfBirth = timestamppb.New(customer.DateOfBirth)
		customerPb.MemberCode = customer.MemberCode
		customerPb.CreatedAt = timestamppb.New(customer.CreatedAt)
		customerPb.UpdatedAt = timestamppb.New(customer.UpdatedAt)

		customerPbRes = append(customerPbRes, customerPb)
	}

	pRes := &pb.ListCustomerResponse{
		CustomerList: customerPbRes,
		Total:        pg.Total,
		Page:         pg.Page,
	}
	return pRes, nil
}

func (h *CustomerHandler) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*empty.Empty, error) {
	err := h.customerRepo.DeleteCustomer(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, err
	}
	return nil, nil
}

func (h *CustomerHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Customer, error) {
	// get metadata from client
	md, _ := metadata.FromIncomingContext(ctx)

	userEmail := md["user"][0]

	// get user from user service
	u, err := h.userClient.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: userEmail})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// check input password with old password hashed
	if !auth.CheckPasswordHash(req.OldPassword, u.Password) {
		return nil, status.Error(codes.Internal, "wrong old password")
	}

	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	u.Password = hashedPassword
	// update user password
	md = metadata.Pairs("user", "0")
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = h.userClient.UpdateUserPassword(ctx, &pb.UpdateUserPasswordRequest{
		Id:       u.Id,
		Password: u.Password,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get customer associated with user
	c, err := h.customerRepo.GetCustomer(ctx, u.CustomerId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Customer{
		Id:    c.ID,
		Name:  c.Name,
		Email: c.Email,
	}
	return pRes, nil
}
