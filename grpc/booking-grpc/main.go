package main

import (
	"context"
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
	"mock-project/grpc/booking-grpc/handlers"
	"mock-project/grpc/booking-grpc/intercepter"
	"mock-project/grpc/booking-grpc/internal"
	"mock-project/grpc/booking-grpc/repo/dbrepo"
	pb "mock-project/pb/proto"
	"net"
)

func main() {
	err := internal.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	port := viper.GetString("port.booking")
	flightPort := viper.GetString("grpc.flight")
	customerPort := viper.GetString("grpc.customer")
	userPort := viper.GetString("grpc.user")

	flightConn, err := grpc.Dial(flightPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
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
	if err != nil {
		panic(err)
	}

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

	bookingRepo, err := dbrepo.NewPostgresRepo(context.Background())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewCustomerManagerClient(customerConn)
	flightClient := pb.NewFlightManagerClient(flightConn)
	userClient := pb.NewUserManagerClient(userConn)
	handler, err := handlers.NewBookingHandler(bookingRepo, customerClient, flightClient, userClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterBookingManagerServer(server, handler)
	logger.Info(fmt.Sprintf("Listen at port: %v", port))

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
