package repository

import (
	"context"
	"mock-project/grpc/flight-grpc/ent"
	"mock-project/grpc/flight-grpc/request"
	"time"
)

type FlightRepository interface {
	CreateFlight(ctx context.Context, f *ent.Flight) (*ent.Flight, error)
	GetFlight(ctx context.Context, name string) (*ent.Flight, error)
	ListFlight(ctx context.Context, paging request.Paging) ([]*ent.Flight, request.Paging, error)
	UpdateFlight(ctx context.Context, id int64, f *ent.Flight) (*ent.Flight, error)
	DeleteFlight(ctx context.Context, id int64) error
	GetFlightById(ctx context.Context, id int64) (*ent.Flight, error)
	SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time, paging request.Paging) ([]*ent.Flight, request.Paging, error)
}
