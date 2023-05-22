package repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"mock_project/ent"
	"mock_project/ent/flight"
	"mock_project/pb"
	"os"
)

type FlightRepository interface {
	CloseDB()
	GetFlightById(ctx context.Context, model *pb.GetFlightByIdRequest) (*ent.Flight, error)
	CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*ent.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*ent.Flight, error)
	DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error)
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (FlightRepository, error) {
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

func (r *PostgresDB) GetFlightById(ctx context.Context, model *pb.GetFlightByIdRequest) (*ent.Flight, error) {
	res, err := r.Client.Flight.Query().Where(flight.ID(uuid.MustParse(model.Id))).Only(ctx)
	log.Println(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*ent.Flight, error) {
	res, err := r.Client.Flight.Create().SetFlightCode(model.FlightInput.FlightCode).SetFrom(model.FlightInput.From).SetTo(model.FlightInput.To).SetDepartureDate(model.FlightInput.DepartureDate.AsTime()).SetDepartureTime(model.FlightInput.DepartureTime.AsTime()).SetDuration(int(model.FlightInput.Duration)).SetCapacity(int(model.FlightInput.Capacity)).SetAvailableSeat(int(model.FlightInput.AvailableSeat)).SetStatus(flight.Status(model.FlightInput.Status.String())).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (r *PostgresDB) UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*ent.Flight, error) {
	myFlight, err := r.Client.Flight.Query().Where(flight.FlightCode(model.FlightInput.FlightCode)).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired flight", err)
		return nil, err
	}
	myFlight, err = myFlight.Update().SetFlightCode(model.FlightInput.FlightCode).SetFrom(model.FlightInput.From).SetTo(model.FlightInput.To).SetDepartureDate(model.FlightInput.DepartureDate.AsTime()).SetDepartureTime(model.FlightInput.DepartureTime.AsTime()).SetDuration(int(model.FlightInput.Duration)).SetCapacity(int(model.FlightInput.Capacity)).SetAvailableSeat(int(model.FlightInput.AvailableSeat)).SetStatus(flight.Status(model.FlightInput.Status.String())).Save(ctx)
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	return myFlight, nil
}

func (r *PostgresDB) DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error) {
	err := r.Client.Flight.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
