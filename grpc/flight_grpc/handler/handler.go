package handler

import (
	"context"
	"log"
	_ "mock_project/ent/flight"
	"mock_project/grpc/flight_grpc/repository"
	"mock_project/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightServiceImp interface {
	GetFlightByID(ctx context.Context, model *pb.GetFlightByIdRequest) (*pb.Flight, error)
	CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*pb.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*pb.Flight, error)
	DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error)
}

type FlightService struct {
	FlightRepository repository.FlightRepository
	pb.UnimplementedFlightServiceRPCServer
}

func NewFlightrHandler(FlightRepository repository.FlightRepository) (*FlightService, error) {
	return &FlightService{
		FlightRepository: FlightRepository,
	}, nil
}

func (s *FlightService) GetFlightByID(ctx context.Context, model *pb.GetFlightByIdRequest) (*pb.Flight, error) {
	res, err := s.FlightRepository.GetFlightById(ctx, model)
	if err != nil {
		return nil, err
	}

	return &pb.Flight{
		Id:            res.ID.String(),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: timestamppb.New(res.DepartureDate),
		DepartureTime: timestamppb.New(res.DepartureTime),
		Duration:      int32(res.Duration),
		Capacity:      int32(res.Capacity),
		AvailableSeat: int32(res.AvailableSeat),
		Status:        pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}

func (s *FlightService) CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*pb.Flight, error) {
	res, err := s.FlightRepository.CreateFlight(ctx, model)
	if err != nil {
		log.Println("Created failed:", err)
		return nil, err
	}

	return &pb.Flight{
		Id:            res.ID.String(),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: timestamppb.New(res.DepartureDate),
		DepartureTime: timestamppb.New(res.DepartureTime),
		Duration:      int32(res.Duration),
		Capacity:      int32(res.Capacity),
		AvailableSeat: int32(res.AvailableSeat),
		Status:        pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}


func (s *FlightService) UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*pb.Flight, error) {
	res, err := s.FlightRepository.UpdateFlightByFlightCode(ctx, model)
	if err != nil {
		log.Fatalf("Updated failed: %s", err)
	}
	return &pb.Flight{
		Id:            res.ID.String(),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: timestamppb.New(res.DepartureDate),
		DepartureTime: timestamppb.New(res.DepartureTime),
		Duration:      int32(res.Duration),
		Capacity:      int32(res.Capacity),
		AvailableSeat: int32(res.AvailableSeat),
		Status:        pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}


func (s *FlightService) DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (*pb.DeleteFlightResponse, error) {
	res, err := s.FlightRepository.DeleteFlight(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteFlightResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteFlightResponse{
		IsDeleted: res,
	}, nil
}
