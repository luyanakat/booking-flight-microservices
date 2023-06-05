package dbrepo

import (
	"context"
	"mock-project/grpc/flight-grpc/ent"
	"mock-project/grpc/flight-grpc/ent/flight"
	"mock-project/grpc/flight-grpc/request"
	"time"
)

func (m postgresDBRepo) CreateFlight(ctx context.Context, f *ent.Flight) (*ent.Flight, error) {
	return m.client.Flight.Create().
		SetName(f.Name).
		SetFrom(f.From).
		SetTo(f.To).
		SetDepartureDate(f.DepartureDate).
		SetArrivalDate(f.ArrivalDate).
		SetAvailableFirstSlot(f.AvailableFirstSlot).
		SetAvailableEconomySlot(f.AvailableEconomySlot).
		SetStatus(f.Status).
		Save(ctx)
}

func (m postgresDBRepo) UpdateFlight(ctx context.Context, id int64, f *ent.Flight) (*ent.Flight, error) {
	return m.client.Flight.UpdateOneID(id).
		SetName(f.Name).
		SetFrom(f.From).
		SetTo(f.To).
		SetDepartureDate(f.DepartureDate).
		SetArrivalDate(f.ArrivalDate).
		SetAvailableFirstSlot(f.AvailableFirstSlot).
		SetAvailableEconomySlot(f.AvailableEconomySlot).
		SetStatus(f.Status).
		Save(ctx)
}

func (m postgresDBRepo) GetFlight(ctx context.Context, name string) (*ent.Flight, error) {
	return m.client.Flight.Query().Where(flight.Name(name)).Only(ctx)
}

func (m postgresDBRepo) GetFlightById(ctx context.Context, id int64) (*ent.Flight, error) {
	return m.client.Flight.Query().Where(flight.ID(id)).Only(ctx)
}

func (m postgresDBRepo) ListFlight(ctx context.Context, paging request.Paging) ([]*ent.Flight, request.Paging, error) {
	total, err := m.client.Flight.Query().Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	flights, err := m.client.Flight.Query().Offset(int((paging.Page - 1) * paging.Limit)).Limit(int(paging.Limit)).All(ctx)
	if err != nil {
		return nil, paging, err
	}
	return flights, paging, nil
}

func (m postgresDBRepo) DeleteFlight(ctx context.Context, id int64) error {
	return m.client.Flight.DeleteOneID(id).Exec(ctx)
}

func (m postgresDBRepo) SearchFlight(ctx context.Context, from, to string, departureDate, arrivalDate time.Time, paging request.Paging) ([]*ent.Flight, request.Paging, error) {
	total, err := m.client.Flight.Query().Where(
		flight.DepartureDateGTE(departureDate),
		flight.ArrivalDateLTE(arrivalDate),
	).Count(ctx)
	if err != nil {
		return nil, paging, err
	}

	paging.Total = int64(total)
	paging.Process()

	flights, err := m.client.Flight.Query().
		Where(
			flight.From(from),
			flight.To(to),
			flight.DepartureDateGTE(departureDate),
			flight.ArrivalDateLTE(arrivalDate),
			flight.StatusEQ(flight.StatusAvailable),
		).
		Offset(int((paging.Page - 1) * paging.Limit)).
		Limit(int(paging.Limit)).
		All(ctx)

	if err != nil {
		return nil, paging, err
	}
	return flights, paging, nil
}
