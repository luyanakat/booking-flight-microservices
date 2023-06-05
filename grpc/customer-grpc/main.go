package main

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"mock-project/grpc/customer-grpc/handlers"
	"mock-project/grpc/customer-grpc/intercepter"
	internal "mock-project/grpc/customer-grpc/internal/config"
	"mock-project/grpc/customer-grpc/repo/dbrepo"
	pb "mock-project/pb/proto"
	"net"
)

func main() {
	err := internal.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}
	port := viper.GetString("port.customer")
	userPort := viper.GetString("grpc.user")

	userConn, err := grpc.Dial(userPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
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
	if err != nil {
		panic(err)
	}

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
		)),
	)

	customerRepo, err := dbrepo.NewPostgresRepo()
	if err != nil {
		panic(err)
	}

	userClient := pb.NewUserManagerClient(userConn)
	handler, err := handlers.NewCustomerHandler(customerRepo, userClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterCustomerManagerServer(server, handler)
	logger.Info(fmt.Sprintf("Listen at port: %v", port))

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
