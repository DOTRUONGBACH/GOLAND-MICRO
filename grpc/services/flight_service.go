package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/flight"
	"mock_project/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightServiceClient interface {
	GetFlightById(ctx context.Context, id string) (*ent.Flight, error)
	CreateFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error)
	DeleteFlight(ctx context.Context,  id string) (bool, error)
}

type FlightHandler struct {
	FlightClient pb.FlightServiceRPCClient
}

func NewFlightHandler(FlightClient pb.FlightServiceRPCClient) FlightServiceClient {
	return &FlightHandler{
		FlightClient: FlightClient,
	}
}

func (h *FlightHandler) CreateFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error) {
	res, err := h.FlightClient.CreateFlight(ctx, &pb.CreateFlightRequest{FlightInput: &pb.FlightInput{FlightCode: input.FlightInput.FlightCode,
		From: input.FlightInput.From,
		To:   input.FlightInput.To, DepartureDate: timestamppb.New(input.FlightInput.DepartureDate),
		DepartureTime: timestamppb.New(input.FlightInput.DepartureTime),
		Duration:      int32(input.FlightInput.Duration),
		Capacity:      int32(input.FlightInput.Capacity),
		AvailableSeat: int32(input.FlightInput.AvailableSeat),
		Status:        pb.FlightInput_Status(pb.FlightInput_Status_value[input.FlightInput.Status.String()])}})

	if err != nil {
		return nil, err
	}

	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.FlightCode,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		AvailableSeat: int(res.AvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) GetFlightById(ctx context.Context, id string) (*ent.Flight, error) {
	res, err := h.FlightClient.GetFlightById(ctx, &pb.GetFlightByIdRequest{Id: id})
	if err != nil {
		log.Println("Can not find desired flight", err)
		return nil, err
	}
	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.FlightCode,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		AvailableSeat: int(res.AvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) UpdateFlightByFlightCode(ctx context.Context, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error) {
	res, err := h.FlightClient.UpdateFlightByFlightCode(ctx, &pb.UpdateFlightByFlightCodeRequest{FlightInput: &pb.FlightInput{FlightCode: input.FlightInput.FlightCode,
		From: input.FlightInput.From,
		To:   input.FlightInput.To, DepartureDate: timestamppb.New(input.FlightInput.DepartureDate),
		DepartureTime: timestamppb.New(input.FlightInput.DepartureTime),
		Duration:      int32(input.FlightInput.Duration),
		Capacity:      int32(input.FlightInput.Capacity),
		AvailableSeat: int32(input.FlightInput.AvailableSeat),
		Status:        pb.FlightInput_Status(pb.FlightInput_Status_value[input.FlightInput.Status.String()])}})
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}

	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.FlightCode,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		AvailableSeat: int(res.AvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil

}

func (h *FlightHandler) DeleteFlight(ctx context.Context, id string) (bool, error) {
	res, err := h.FlightClient.DeleteFlight(ctx, &pb.DeleteFlightRequest{Id: id})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}
