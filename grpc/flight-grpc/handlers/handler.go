package handlers

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mock-project/grpc/flight-grpc/ent"
	"mock-project/grpc/flight-grpc/ent/flight"
	"mock-project/grpc/flight-grpc/internal"
	repository "mock-project/grpc/flight-grpc/repo"
	"mock-project/grpc/flight-grpc/request"
	pb "mock-project/pb/proto"
)

type FlightHandler struct {
	pb.UnimplementedFlightManagerServer
	flightRepo    repository.FlightRepository
	bookingClient pb.BookingManagerClient
}

func NewFlightHandler(flightRepo repository.FlightRepository, bookingClient pb.BookingManagerClient) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepo:    flightRepo,
		bookingClient: bookingClient,
	}, nil
}

func (h *FlightHandler) SearchFlight(ctx context.Context, req *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}

	// repo
	list, pg, err := h.flightRepo.SearchFlight(ctx, req.From, req.To, req.DepartureDate.AsTime(), req.ArrivalDate.AsTime(), paging)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// append ent type to protobuf response
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.ID
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableFirstSlot = int64(fl.AvailableFirstSlot)
		flightPb.AvailableEconomySlot = int64(fl.AvailableEconomySlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.SearchFlightResponse{
		FlightList: flightPbRes,
		Total:      pg.Total,
		Page:       pg.Page,
	}
	return pRes, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	statusParse, _ := internal.ParseString(req.Status)

	fl := &ent.Flight{
		Name:                 req.Name,
		From:                 req.From,
		To:                   req.To,
		DepartureDate:        req.DepartureDate.AsTime(),
		ArrivalDate:          req.ArrivalDate.AsTime(),
		AvailableFirstSlot:   int(req.AvailableFirstSlot),
		AvailableEconomySlot: int(req.AvailableEconomySlot),
		Status:               statusParse,
	}

	f, err := h.flightRepo.CreateFlight(ctx, fl)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Flight{}
	if err := copier.Copy(&pRes, f); err != nil {
		return nil, err
	}
	pRes.Id = f.ID

	return pRes, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	statusParse, _ := internal.ParseString(req.Status)

	flReq := &ent.Flight{
		Name:                 req.Name,
		DepartureDate:        req.DepartureDate.AsTime(),
		ArrivalDate:          req.ArrivalDate.AsTime(),
		AvailableFirstSlot:   int(req.AvailableFirstSlot),
		AvailableEconomySlot: int(req.AvailableEconomySlot),
		Status:               statusParse,
	}

	flightInput := internal.CheckFlightEmptyInput(flReq, fl)
	if fl.Status == flight.StatusCancel {
		return nil, status.Error(codes.InvalidArgument, "flight was canceled, can't update status as cancel")
	} else if fl.Status == flight.StatusArrived {
		return nil, status.Error(codes.InvalidArgument, "flight was arrived, can't update status as arrived")
	}

	f, err := h.flightRepo.UpdateFlight(ctx, fl.ID, flightInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if f.Status == flight.StatusCancel {
		_, err := h.bookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{FlightId: f.ID, Status: flight.StatusCancel.String()})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else if f.Status == flight.StatusArrived {
		_, err := h.bookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{FlightId: f.ID, Status: flight.StatusArrived.String()})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	pRes := &pb.Flight{}
	if err := copier.Copy(&pRes, f); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pRes.Id = f.ID
	pRes.DepartureDate = timestamppb.New(f.DepartureDate)
	pRes.ArrivalDate = timestamppb.New(f.ArrivalDate)

	return pRes, nil
}

func (h *FlightHandler) GetFlight(ctx context.Context, req *pb.GetFlightRequest) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlight(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Flight{
		Id:                   fl.ID,
		Name:                 fl.Name,
		From:                 fl.From,
		To:                   fl.To,
		DepartureDate:        timestamppb.New(fl.DepartureDate),
		ArrivalDate:          timestamppb.New(fl.ArrivalDate),
		AvailableFirstSlot:   int64(fl.AvailableFirstSlot),
		AvailableEconomySlot: int64(fl.AvailableEconomySlot),
		Status:               fl.Status.String(),
		CreatedAt:            timestamppb.New(fl.CreatedAt),
		UpdatedAt:            timestamppb.New(fl.UpdatedAt),
	}

	return pRes, nil
}

func (h *FlightHandler) GetFlightById(ctx context.Context, req *pb.GetFlightRequest) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Flight{
		Id:                   fl.ID,
		Name:                 fl.Name,
		From:                 fl.From,
		To:                   fl.To,
		DepartureDate:        timestamppb.New(fl.DepartureDate),
		ArrivalDate:          timestamppb.New(fl.ArrivalDate),
		AvailableFirstSlot:   int64(fl.AvailableFirstSlot),
		AvailableEconomySlot: int64(fl.AvailableEconomySlot),
		Status:               fl.Status.String(),
		CreatedAt:            timestamppb.New(fl.CreatedAt),
		UpdatedAt:            timestamppb.New(fl.UpdatedAt),
	}
	return pRes, nil
}

func (h *FlightHandler) ListFlight(ctx context.Context, req *pb.ListFlightRequest) (*pb.ListFlightResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}

	list, pg, err := h.flightRepo.ListFlight(ctx, paging)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// append ent type to protobuf response
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.ID
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableFirstSlot = int64(fl.AvailableFirstSlot)
		flightPb.AvailableEconomySlot = int64(fl.AvailableEconomySlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.ListFlightResponse{
		FlightList: flightPbRes,
		Total:      pg.Total,
		Page:       pg.Page,
	}
	return pRes, nil
}

func (h *FlightHandler) UpdateFlightSlot(ctx context.Context, req *pb.UpdateFlightSlotRequest) (*pb.Flight, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	// Reduce slot if calling from create booking
	if md["update"][0] == "create" {
		fl, _ := h.flightRepo.GetFlightById(ctx, req.Id)
		if req.TicketType == 2 {
			fl.AvailableEconomySlot = fl.AvailableEconomySlot - 1
		} else {
			fl.AvailableFirstSlot = fl.AvailableFirstSlot - 1
		}

		updatedFlight, err := h.flightRepo.UpdateFlight(ctx, req.Id, fl)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pRes := &pb.Flight{Id: updatedFlight.ID}
		return pRes, nil
	}

	// Add slot if calling from cancel booking
	if md["update"][0] == "delete" {
		fl, _ := h.flightRepo.GetFlightById(ctx, req.Id)
		if req.TicketType == 2 {
			fl.AvailableEconomySlot = fl.AvailableEconomySlot + 1
		} else {
			fl.AvailableFirstSlot = fl.AvailableFirstSlot + 1
		}

		updatedFlight, err := h.flightRepo.UpdateFlight(ctx, req.Id, fl)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pRes := &pb.Flight{Id: updatedFlight.ID}
		return pRes, nil
	}

	return nil, nil
}

func (h *FlightHandler) DeleteFlight(ctx context.Context, req *pb.DeleteFlightRequest) (*empty.Empty, error) {
	err := h.flightRepo.DeleteFlight(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, err
	}
	return nil, nil
}
