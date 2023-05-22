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
func (r *accountResolver) ID(ctx context.Context, obj *ent.Account) (string, error) {
	return obj.ID.String(), nil
}

// Role is the resolver for the role field.
func (r *accountResolver) Role(ctx context.Context, obj *ent.Account) (ent.Role, error) {
	return ent.Role(obj.Role), nil
}

// Role is the resolver for the role field.
func (r *updateAccountInputResolver) Role(ctx context.Context, obj *ent.UpdateAccountInput, data ent.Role) error {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Customer is the resolver for the customer field.
func (r *updateAccountInputResolver) Customer(ctx context.Context, obj *ent.UpdateAccountInput, data *ent.CustomerInput) error {
	panic(fmt.Errorf("not implemented: Customer - customer"))
}

// Account returns graphql1.AccountResolver implementation.
func (r *Resolver) Account() graphql1.AccountResolver { return &accountResolver{r} }

// UpdateAccountInput returns graphql1.UpdateAccountInputResolver implementation.
func (r *Resolver) UpdateAccountInput() graphql1.UpdateAccountInputResolver {
	return &updateAccountInputResolver{r}
}

type accountResolver struct{ *Resolver }
type updateAccountInputResolver struct{ *Resolver }
