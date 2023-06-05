package main

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"mock-project/grpc/flight-grpc/handlers"
	"mock-project/grpc/flight-grpc/intercepter"
	"mock-project/grpc/flight-grpc/internal"
	"mock-project/grpc/flight-grpc/repo/dbrepo"
	pb "mock-project/pb/proto"
	"net"
)

func main() {
	err := internal.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}
	port := viper.GetString("port.flight")
	bookingPort := viper.GetString("grpc.booking")

	bookingConn, err := grpc.Dial(bookingPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
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

	flightRepo, err := dbrepo.NewPostgresRepo()
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBookingManagerClient(bookingConn)
	handler, err := handlers.NewFlightHandler(flightRepo, bookingClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterFlightManagerServer(server, handler)
	logger.Info(fmt.Sprintf("Listen at port: %v", port))

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
