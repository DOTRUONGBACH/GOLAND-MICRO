package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mock_project/ent"
	graphql1 "mock_project/graphql"

	"github.com/google/uuid"
)

// DeleteCustomer is the resolver for the deleteCustomer field.
func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (bool, error) {
	return r.customerService.DeleteCustomer(ctx, uuid.MustParse(id))
}

// NewAccount is the resolver for the newAccount field.
func (r *mutationResolver) NewAccount(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error) {
	return r.accountService.CreateAccount(ctx, input)
}

// AccResetPassword is the resolver for the accResetPassword field.
func (r *mutationResolver) AccResetPassword(ctx context.Context, input ent.AccountResetPasswordInput) (string, error) {
	return r.accountService.AccountResetPassword(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input ent.AccountLogin) (*ent.AccountLoginResponse, error) {
	res, err := r.accountService.Login(ctx, input)
	return &ent.AccountLoginResponse{
		Token:  res.Token,
		Status: res.Status,
	}, err
}

// NewFlight is the resolver for the newFlight field.
func (r *mutationResolver) NewFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error) {
	return r.flightService.CreateFlight(ctx, input)
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
