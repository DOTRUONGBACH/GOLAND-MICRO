package repository
import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/booking"
	"mock_project/pb"
	"os"
	_ "github.com/lib/pq"
	"github.com/google/uuid"
)

type BookingRepository interface {
	GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*ent.Booking, error)
	CreateBooking(ctx context.Context, model *pb.CreateBookingRequest) (*ent.Booking, error)
	UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*ent.Booking, error)	
	DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool ,error)
	CloseDB()
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (BookingRepository, error) {
	client, err := ent.Open("postgres", connection_str)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating db schema from the automation migration tool")
		os.Exit(1)
	}

	return &PostgresDB{Client: client}, nil
}


func (r *PostgresDB) CloseDB() {
	r.Client.Close()
}

func (r *PostgresDB) GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*ent.Booking, error){
	res, err := r.Client.Booking.Get(ctx,uuid.MustParse(model.Id))
	if err != nil {
		log.Println("Can not find the desired boooking",err)
		return nil, err
	}

	return res, nil
}

func (r *PostgresDB) CreateBooking(ctx context.Context, model *pb.CreateBookingRequest) (*ent.Booking, error){

	totalCost := float64(model.TotalTicket*100)


	res, err := r.Client.Booking.Create().SetTotalCost(totalCost).SetTotalTicket(int(model.TotalTicket)).SetStatus(booking.Status(model.Status.String())).Save(ctx)

	if err != nil {
		log.Println("Create boooking failed",err)
		return nil, err
	}

	return res, nil
}


func (r *PostgresDB) UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*ent.Booking, error) {
	res, err1 := r.Client.Booking.Get(ctx, uuid.MustParse(model.Id))
	if err1 != nil {
		log.Println("Can not find the desired boooking", err1)
		return nil, err1
	}

	res, err2 := res.Update().SetStatus(booking.Status(model.Status)).Save(ctx)

	if err2 != nil {
		log.Println("Can not update the boooking status", err2)
		return nil, err2
	}

	return res, nil
}

func (r *PostgresDB) DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool, error) {
	err := r.Client.Booking.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

