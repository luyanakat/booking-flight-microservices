package handlers

import (
	"context"
	"database/sql"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mock-project/grpc/user-grpc/ent"
	"mock-project/grpc/user-grpc/internal/auth"
	repository "mock-project/grpc/user-grpc/repo"
	"mock-project/grpc/user-grpc/request"
	pb "mock-project/pb/proto"
	"strconv"
)

type UserHandler struct {
	pb.UnimplementedUserManagerServer
	customerClient pb.CustomerManagerClient
	userRepo       repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository, customerClient pb.CustomerManagerClient) (*UserHandler, error) {
	return &UserHandler{
		userRepo:       userRepo,
		customerClient: customerClient,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := h.userRepo.GetUserByCustomerId(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.User{
		Id:         user.ID,
		Email:      user.Email,
		Password:   user.Password,
		CustomerId: user.CustomerID,
		AccessId:   user.AccessID,
	}
	return pRes, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &ent.User{}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user.Password = hashedPassword
	user.Email = req.Email
	user.AccessID = req.AccessId

	_, err = h.userRepo.GetAccessLevel(ctx, user.AccessID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "access level not found")
		}
		return nil, err
	}

	u, err := h.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.User{}
	if err := copier.Copy(&pRes, u); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pRes.Id = u.ID
	return pRes, nil
}

func (h *UserHandler) RegisterCustomer(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	uE, _ := h.userRepo.GetUserByEmail(ctx, req.Email)
	// if customer wasn't created account
	if uE == nil {
		// register new customer
		customer, err := h.customerClient.CreateCustomer(ctx, &pb.Customer{
			Name:           req.CustomerName,
			Email:          req.Email,
			Address:        req.Address,
			PhoneNumber:    req.PhoneNumber,
			IdentifyNumber: req.IdentifyNumber,
			DateOfBirth:    req.DateOfBirth,
			MemberCode:     req.MemberCode,
		})
		if err != nil {
			return nil, err
		}

		if req.AccessId == 0 {
			req.AccessId = 2
		}

		passHashed, err := auth.HashPassword(req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// create user with customer id was generated before
		user := &ent.User{
			Email:      req.Email,
			Password:   passHashed,
			CustomerID: customer.Id,
			AccessID:   req.AccessId,
		}

		u, err := h.userRepo.CreateUser(ctx, user)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// gen jwt token
		token, err := auth.GenerateToken(req.Email)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		pRes := &pb.RegisterResponse{
			UserId:     u.ID,
			CustomerId: customer.Id,
			JwtToken:   token,
		}

		return pRes, nil
	} else { // if customer was created account
		return nil, status.Error(codes.AlreadyExists, "email have been registered")
	}
}

func (h *UserHandler) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}

	list, pg, err := h.userRepo.ListUser(ctx, paging)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var userPbRes []*pb.User
	for _, user := range list {
		userPb := &pb.User{}
		userPb.Id = user.ID
		userPb.Email = user.Email
		userPb.CustomerId = user.CustomerID
		userPb.AccessId = user.AccessID
		userPb.CreatedAt = timestamppb.New(user.CreatedAt)
		userPb.UpdatedAt = timestamppb.New(user.UpdatedAt)

		userPbRes = append(userPbRes, userPb)
	}

	pRes := &pb.ListUserResponse{
		UserList: userPbRes,
		Total:    pg.Total,
		Page:     pg.Page,
	}
	return pRes, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	idConV, _ := strconv.Atoi(md["user"][0])
	u, err := h.userRepo.GetUser(ctx, int64(idConV))
	if err != nil {
		return nil, err
	}

	if req.AccessId != 0 {
		u.AccessID = req.AccessId
	}
	userGet, err := h.userRepo.UpdateUser(ctx, u.ID, u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.User{
		Id:         userGet.ID,
		Email:      userGet.Email,
		Password:   userGet.Password,
		CustomerId: userGet.CustomerID,
		AccessId:   userGet.AccessID,
	}
	return pRes, nil
}

func (h *UserHandler) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.User, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	// calling from update user password
	if md["user"][0] != "0" {
		idConV, _ := strconv.Atoi(md["user"][0])
		u, err := h.userRepo.GetUser(ctx, int64(idConV))
		if err != nil {
			return nil, err
		}
		if !auth.CheckPasswordHash(req.OldPassword, u.Password) {
			return nil, status.Error(codes.InvalidArgument, "old password is not correct")
		}

		passHashed, err := auth.HashPassword(req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		userUpdated, err := h.userRepo.UpdateUserPassword(ctx, u.ID, passHashed)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &pb.User{Id: userUpdated.ID}, nil
	}

	// calling from update customer password
	if md["user"][0] == "0" {
		u, err := h.userRepo.UpdateUserPassword(ctx, req.Id, req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &pb.User{Id: u.ID}, nil
	}
	return nil, nil
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.User, error) {
	u, err := h.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
	}

	pRes := &pb.User{
		Id:         u.ID,
		Email:      u.Email,
		CustomerId: u.CustomerID,
		AccessId:   u.AccessID,
		Password:   u.Password,
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdatedAt:  timestamppb.New(u.UpdatedAt),
	}
	return pRes, nil
}

func (h *UserHandler) Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, err := h.userRepo.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// check password
	if !auth.CheckPasswordHash(loginRequest.Password, u.Password) {
		return nil, status.Error(codes.Unauthenticated, "wrong username or password")
	}

	token, err := auth.GenerateToken(loginRequest.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.LoginResponse{
		Id:    u.ID,
		Token: token,
	}
	return pRes, nil
}

func (h *UserHandler) ParseToken(ctx context.Context, tokenRequest *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	email, err := auth.ParseToken(tokenRequest.Token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	u, err := h.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pUser := &pb.User{
		Id:         u.ID,
		Email:      u.Email,
		CustomerId: u.CustomerID,
		AccessId:   u.AccessID,
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdatedAt:  timestamppb.New(u.UpdatedAt),
	}
	pRes := &pb.ParseTokenResponse{
		User: pUser,
	}
	return pRes, nil
}

func (h *UserHandler) GetAccessLevel(ctx context.Context, req *pb.GetAccessLevelRequest) (*pb.AccessLevel, error) {
	accessLevel, err := h.userRepo.GetAccessLevel(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.AccessLevel{
		Id:        accessLevel.ID,
		Name:      accessLevel.Name,
		CreatedAt: timestamppb.New(accessLevel.CreatedAt),
		UpdatedAt: timestamppb.New(accessLevel.UpdatedAt),
	}
	return pRes, nil
}
