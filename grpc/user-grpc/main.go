package main

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/spf13/viper"
	"mock-project/grpc/user-grpc/intercepter"
	"mock-project/helper"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"mock-project/grpc/user-grpc/handlers"
	"mock-project/grpc/user-grpc/repo/dbrepo"
	pb "mock-project/pb/proto"
	"net"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	port := viper.GetString("port.user")
	customerPort := viper.GetString("grpc.customer")

	customerConn, err := grpc.Dial(customerPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpcMiddleware.ChainUnaryClient(
			grpcRetry.UnaryClientInterceptor(
				grpcRetry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpcRetry.WithMax(2)),
		)),
		grpc.WithChainStreamInterceptor(grpcMiddleware.ChainStreamClient(
			grpcRetry.StreamClientInterceptor(
				grpcRetry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpcRetry.WithMax(2)),
		)))

	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcRecovery.StreamServerInterceptor(),
		)))

	userRepo, err := dbrepo.NewPostgresRepo(context.Background())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewCustomerManagerClient(customerConn)
	handler, err := handlers.NewUserHandler(userRepo, customerClient)
	if err != nil {
		panic(err)
	}
	reflection.Register(server)
	pb.RegisterUserManagerServer(server, handler)
	logger.Info(fmt.Sprintf("Listen at port: %v", port))

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
