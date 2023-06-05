package repository

import (
	"context"
	"mock-project/grpc/booking-grpc/ent"
	"mock-project/grpc/booking-grpc/request"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, b *ent.Booking) (*ent.Booking, error)
	GetTicketClass(ctx context.Context, id int64) (*ent.Ticket, error)
	UpdateBooking(ctx context.Context, id int64, b *ent.Booking) (*ent.Booking, error)
	GetBookingByCode(ctx context.Context, code string) (*ent.Booking, error)
	GetBookingHistory(ctx context.Context, paging request.Paging, customerId int64) ([]*ent.Booking, request.Paging, error)
	GetBookingByFlight(ctx context.Context, flightId int64) ([]*ent.Booking, error)
	UpdateBookingStatus(ctx context.Context, id int64, status string) error
	ListBooking(ctx context.Context, paging request.Paging) ([]*ent.Booking, request.Paging, error)
}
