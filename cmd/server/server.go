package api

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mock-project/graphql/resolver"
	"mock-project/helper"
	"mock-project/middleware"
	pb "mock-project/pb/proto"
)

func init() {
	err := helper.AutoBindConfig("config.yaml")
	if err != nil {
		panic(err)
	}
}

func getPort(key string) string {
	return viper.GetString(fmt.Sprintf("grpc.%s", key))
}

func Server() *gin.Engine {
	customerConn, err := grpc.Dial(getPort("customer"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	flightConn, err := grpc.Dial(getPort("flight"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userConn, err := grpc.Dial(getPort("user"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	bookingConn, err := grpc.Dial(getPort("booking"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// connect to proto client
	userClient := pb.NewUserManagerClient(userConn)
	customerClient := pb.NewCustomerManagerClient(customerConn)
	flightClient := pb.NewFlightManagerClient(flightConn)
	bookingClient := pb.NewBookingManagerClient(bookingConn)

	// graphQL server
	resolverHandler := handler.NewDefaultServer(resolver.NewSchema(userClient, customerClient, flightClient, bookingClient))
	playGroundHandler := playground.Handler("GraphQL", "/graphql")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.HandleMethodNotAllowed = true

	// use middleware
	r.Use(
		middleware.Auth(userClient),
		middleware.LoggingMiddleware(logger),
		middleware.CorsMiddleware(),
	)

	// Create new GraphQL
	r.POST("/graphql", func(c *gin.Context) {
		resolverHandler.ServeHTTP(c.Writer, c.Request)
	})

	r.OPTIONS("/graphql", func(c *gin.Context) {
		c.Status(200)
	})

	// Enable playground for query/testing
	r.GET("/", func(c *gin.Context) {
		playGroundHandler.ServeHTTP(c.Writer, c.Request)
	})

	return r
}
