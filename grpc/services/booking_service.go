package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/booking"
	_ "mock_project/ent/booking"
	"mock_project/pb"

	"github.com/google/uuid"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

type BookingServiceClient interface {
	// GetBookingByID(ctx context.Context, id string) (*ent.Booking, error)
	CreateBooking(ctx context.Context, input ent.NewBookingInput) (*ent.Booking, error)
	// UpdateBookingStatus(ctx context.Context,input ent.UpdateBookingStatusInput) (*ent.Booking, error)	
	// DeleteBooking(ctx context.Context, id string) (bool ,error)
}

type BookingHandler struct {
	BookingClient pb.BookingServiceRPCClient
	FlightClient pb.FlightServiceRPCClient
}

func NewBookingHandler(BookingClient pb.BookingServiceRPCClient, FlightClient pb.FlightServiceRPCClient) BookingServiceClient {
	return &BookingHandler{
		BookingClient: BookingClient,
	}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, input ent.NewBookingInput) (*ent.Booking, error) {
	myFlight, err := h.FlightClient.GetFlightByFlightCode(ctx, &pb.GetFlightByFlightCodeRequest{FlightCode: input.FlightCode})

	if err != nil {
		log.Println("Flight does not exist",err)
		return nil, err
	}
	if myFlight.AvailableSeat >= int32(input.TotalTicket) && myFlight.Status.String() == "Scheduled"{
		res, err := h.BookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{TotalTicket: int32(input.TotalTicket), Status: pb.CreateBookingRequest_Success})
		if err != nil {
			return nil, err
		}
		return &ent.Booking{
			ID: uuid.MustParse(res.Id),
			TotalCost: float64(res.TotalCost),
			TotalTicket: int(res.TotalTicket),
			Status: booking.StatusSuccess,
		},nil
	}else{
		res, err := h.BookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{TotalTicket: int32(input.TotalTicket), Status: pb.CreateBookingRequest_Fail})
		if err != nil {
			return nil, err
		}
		return &ent.Booking{
			ID: uuid.MustParse(res.Id),
			TotalCost: float64(res.TotalCost),
			TotalTicket: int(res.TotalTicket),
			Status: booking.StatusFail,
		},nil
	}
}

// func (h *BookingHandler) GetBookingByID(ctx context.Context, id string) (*ent.Booking, error) {
// 	// res, err := h.BookingClient.GetBookingByID(ctx, &pb.GetBookingByIdRequest{Id: id})
// 	// if err != nil {
// 	// 	log.Println("Can not find desired Booking", err)
// 	// 	return nil, err
// 	// }
// 	// return &ent.Booking{
// 	// 	ID:            uuid.MustParse(res.Id),
// 	// 	BookingCode:    res.BookingCode,
// 	// 	From:          res.BookingCode,
// 	// 	To:            res.To,
// 	// 	DepartureDate: res.DepartureDate.AsTime(),
// 	// 	DepartureTime: res.DepartureTime.AsTime(),
// 	// 	Duration:      int(res.Duration),
// 	// 	Capacity:      int(res.Capacity),
// 	// 	AvailableSeat: int(res.AvailableSeat),
// 	// 	Status:        Booking.Status(res.Status.String()),
// 	// }, nil
// }

// func (h *BookingHandler) UpdateBookingByBookingCode(ctx context.Context, input ent.UpdateBookingStatusInput) (*ent.Booking, error) {
// 	// res, err := h.BookingClient.UpdateBookingByBookingCode(ctx, &pb.UpdateBookingByBookingCodeRequest{BookingInput: &pb.BookingInput{BookingCode: input.BookingInput.BookingCode,
// 	// 	From: input.BookingInput.From,
// 	// 	To:   input.BookingInput.To, DepartureDate: timestamppb.New(input.BookingInput.DepartureDate),
// 	// 	DepartureTime: timestamppb.New(input.BookingInput.DepartureTime),
// 	// 	Duration:      int32(input.BookingInput.Duration),
// 	// 	Capacity:      int32(input.BookingInput.Capacity),
// 	// 	AvailableSeat: int32(input.BookingInput.AvailableSeat),
// 	// 	Status:        pb.BookingInput_Status(pb.BookingInput_Status_value[input.BookingInput.Status.String()])}})
// 	// if err != nil {
// 	// 	log.Println("Update Booking failed", err)
// 	// 	return nil, err
// 	// }

// 	// return &ent.Booking{
// 	// 	ID:            uuid.MustParse(res.Id),
// 	// 	BookingCode:    res.BookingCode,
// 	// 	From:          res.BookingCode,
// 	// 	To:            res.To,
// 	// 	DepartureDate: res.DepartureDate.AsTime(),
// 	// 	DepartureTime: res.DepartureTime.AsTime(),
// 	// 	Duration:      int(res.Duration),
// 	// 	Capacity:      int(res.Capacity),
// 	// 	AvailableSeat: int(res.AvailableSeat),
// 	// 	Status:        Booking.Status(res.Status.String()),
// 	// }, nil

// }

// func (h *BookingHandler) DeleteBooking(ctx context.Context, id string) (bool, error) {
// 	res, err := h.BookingClient.DeleteBooking(ctx, &pb.DeleteBookingRequest{Id: id})
// 	if err != nil {
// 		return res.IsDeleted, err
// 	}
// 	return res.IsDeleted, nil
// }
