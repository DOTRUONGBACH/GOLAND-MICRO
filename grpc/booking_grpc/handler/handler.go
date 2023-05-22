package handler

import (
	"context"
	"log"
	_ "mock_project/ent/booking"
	"mock_project/grpc/booking_grpc/repository"
	"mock_project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

type BookingServiceImp interface {
	GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*pb.Booking, error)
	CreateBooking(ctx context.Context, model *pb.CreateBookingRequest) (*pb.Booking, error)
	UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*pb.Booking, error)	
	DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool ,error)
}

type BookingService struct {
	BookingRepository repository.BookingRepository
	pb.UnimplementedBookingServiceRPCServer
}

func NewBookingrHandler(BookingRepository repository.BookingRepository) (*BookingService, error) {
	return &BookingService{
		BookingRepository: BookingRepository,
	}, nil
}

func (s *BookingService) GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*pb.Booking, error) {
	res, err := s.BookingRepository.GetBookingByID(ctx, model)
	
	if err != nil {
		log.Println("Can not find the desired boooking",err)
		return nil, err
	}

	return &pb.Booking{
		Id: res.ID.String(),
		TotalCost: float32(res.TotalCost),
		TotalTicket: int32(res.TotalTicket),
		Status: pb.Booking_Status(pb.Booking_Status_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	},nil 
}

func (s *BookingService) CreateBooking(ctx context.Context, model *pb.CreateBookingRequest) (*pb.Booking, error) {
	res, err := s.BookingRepository.CreateBooking(ctx, model)

	if err != nil {
		log.Println("Create boooking failed",err)
		return nil, err
	}

	return &pb.Booking{
		Id: res.ID.String(),
		TotalCost: float32(res.TotalCost),
		TotalTicket: int32(res.TotalTicket),
		Status: pb.Booking_Status(pb.Booking_Status_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	},nil 
}


func (s *BookingService) UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*pb.Booking, error) {
	res, err := s.BookingRepository.UpdateBookingStatus(ctx, model)

	if err != nil {
		log.Println("Update boooking failed",err)
		return nil, err
	}

	return &pb.Booking{
		Id: res.ID.String(),
		TotalCost: float32(res.TotalCost),
		TotalTicket: int32(res.TotalTicket),
		Status: pb.Booking_Status(pb.Booking_Status_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	},nil 
}


func (s *BookingService) DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (*pb.DeleteBookingResponse, error) {
	res, err := s.BookingRepository.DeleteBooking(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteBookingResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteBookingResponse{
		IsDeleted: res,
	}, nil
}
