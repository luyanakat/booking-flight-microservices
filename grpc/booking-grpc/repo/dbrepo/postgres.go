package dbrepo

import (
	"context"
	"mock-project/grpc/booking-grpc/ent"
	"mock-project/grpc/booking-grpc/ent/booking"
	"mock-project/grpc/booking-grpc/ent/ticket"
	"mock-project/grpc/booking-grpc/internal"
	"mock-project/grpc/booking-grpc/request"
)

func (m postgresDBRepo) CreateBooking(ctx context.Context, b *ent.Booking) (*ent.Booking, error) {
	return m.client.Booking.Create().
		SetCustomerID(b.CustomerID).
		SetFlightID(b.FlightID).
		SetStatus(b.Status).
		SetTicketID(b.TicketID).
		Save(ctx)
}

func (m postgresDBRepo) UpdateBooking(ctx context.Context, id int64, b *ent.Booking) (*ent.Booking, error) {
	return m.client.Booking.UpdateOneID(id).
		SetCustomerID(b.CustomerID).
		SetFlightID(b.FlightID).
		SetStatus(b.Status).
		SetTicketID(b.TicketID).
		Save(ctx)
}

func (m postgresDBRepo) GetBookingByCode(ctx context.Context, code string) (*ent.Booking, error) {
	return m.client.Booking.Query().Where(booking.Code(code)).Only(ctx)
}

func (m postgresDBRepo) GetTicketClass(ctx context.Context, id int64) (*ent.Ticket, error) {
	return m.client.Ticket.Query().Where(ticket.ID(id)).Only(ctx)
}

func (m postgresDBRepo) GetBookingHistory(ctx context.Context, paging request.Paging, customerId int64) ([]*ent.Booking, request.Paging, error) {
	total, err := m.client.Booking.Query().Where(booking.CustomerID(customerId)).Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	bookingHistory, err := m.client.Booking.Query().Where(booking.CustomerID(customerId)).Offset(int((paging.Page - 1) * paging.Limit)).Limit(int(paging.Limit)).All(ctx)
	if err != nil {
		return nil, paging, err
	}
	return bookingHistory, paging, nil
}

func (m postgresDBRepo) GetBookingByFlight(ctx context.Context, flightId int64) ([]*ent.Booking, error) {
	return m.client.Booking.Query().Where(booking.FlightID(flightId)).All(ctx)
}

func (m postgresDBRepo) UpdateBookingStatus(ctx context.Context, id int64, status string) error {
	convertedEnum, _ := internal.ParseString(status)
	return m.client.Booking.Update().Where(booking.FlightID(id)).SetStatus(convertedEnum).Exec(ctx)
}

func (m postgresDBRepo) ListBooking(ctx context.Context, paging request.Paging) ([]*ent.Booking, request.Paging, error) {
	total, err := m.client.Booking.Query().Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	bks, err := m.client.Booking.Query().Offset(int((paging.Page - 1) * paging.Limit)).Limit(int(paging.Limit)).All(ctx)
	if err != nil {
		return nil, paging, err
	}
	return bks, paging, nil
}
