package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"mock-project/graphql/graph"
	pb "mock-project/pb/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userClient     pb.UserManagerClient
	customerClient pb.CustomerManagerClient
	flightClient   pb.FlightManagerClient
	bookingClient  pb.BookingManagerClient
}

func NewSchema(userClient pb.UserManagerClient, customerClient pb.CustomerManagerClient, flightClient pb.FlightManagerClient, bookingClient pb.BookingManagerClient) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: &Resolver{
			userClient:     userClient,
			customerClient: customerClient,
			flightClient:   flightClient,
			bookingClient:  bookingClient,
		},
	})
}
