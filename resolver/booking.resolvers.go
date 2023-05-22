package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
)

// ID is the resolver for the id field.
func (r *bookingResolver) ID(ctx context.Context, obj *ent.Booking) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// TotalCost is the resolver for the totalCost field.
func (r *bookingResolver) TotalCost(ctx context.Context, obj *ent.Booking) (string, error) {
	panic(fmt.Errorf("not implemented: TotalCost - totalCost"))
}

// Status is the resolver for the status field.
func (r *bookingResolver) Status(ctx context.Context, obj *ent.Booking) (ent.BookingStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Booking returns graphql1.BookingResolver implementation.
func (r *Resolver) Booking() graphql1.BookingResolver { return &bookingResolver{r} }

type bookingResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *bookingResolver) TotalTicket(ctx context.Context, obj *ent.Booking) (string, error) {
	panic(fmt.Errorf("not implemented: TotalTicket - totalTicket"))
}
