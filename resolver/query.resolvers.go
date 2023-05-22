package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"mock_project/ent"
	graphql1 "mock_project/graphql"

	"github.com/google/uuid"
)

// Customers is the resolver for the Customers field.
func (r *queryResolver) Customers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CustomerOrder) (*ent.CustomerConnection, error) {
	return r.client.Customer.Query().Paginate(ctx, after, first, before, last, ent.WithCustomerOrder(orderBy))
}

// Accounts is the resolver for the Accounts field.
func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder) (*ent.AccountConnection, error) {
	return r.client.Account.Query().Paginate(ctx, after, first, before, last, ent.WithAccountOrder(orderBy))
}

// GetAccountByID is the resolver for the GetAccountByID field.
func (r *queryResolver) GetAccountByID(ctx context.Context, input string) (*ent.AccountByIDResponse, error) {
	res, err := r.accountService.GetAccountByID(ctx, uuid.MustParse(input))
	log.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC", res)
	return &ent.AccountByIDResponse{
		Email:    res.Email,
		Role:     ent.Role(res.Role),
		AccOwner: &ent.Customer{ID: res.Edges.AccOwner.ID, Name: res.Edges.AccOwner.Name, CitizenID: res.Edges.AccOwner.CitizenID, Phone: res.Edges.AccOwner.Phone, Address: res.Edges.AccOwner.Address, Gender: res.Edges.AccOwner.Gender, DateOfBirth: res.Edges.AccOwner.DateOfBirth},
	}, err
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
