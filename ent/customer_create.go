// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock_project/ent/account"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CustomerCreate) SetName(s string) *CustomerCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetCitizenID sets the "citizen_id" field.
func (cc *CustomerCreate) SetCitizenID(s string) *CustomerCreate {
	cc.mutation.SetCitizenID(s)
	return cc
}

// SetPhone sets the "phone" field.
func (cc *CustomerCreate) SetPhone(s string) *CustomerCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *CustomerCreate) SetAddress(s string) *CustomerCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetGender sets the "gender" field.
func (cc *CustomerCreate) SetGender(c customer.Gender) *CustomerCreate {
	cc.mutation.SetGender(c)
	return cc
}

// SetDateOfBirth sets the "date_of_birth" field.
func (cc *CustomerCreate) SetDateOfBirth(t time.Time) *CustomerCreate {
	cc.mutation.SetDateOfBirth(t)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CustomerCreate) SetCreatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CustomerCreate) SetUpdatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CustomerCreate) SetID(u uuid.UUID) *CustomerCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableID(u *uuid.UUID) *CustomerCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (cc *CustomerCreate) AddAccountIDs(ids ...uuid.UUID) *CustomerCreate {
	cc.mutation.AddAccountIDs(ids...)
	return cc
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (cc *CustomerCreate) AddAccounts(a ...*Account) *CustomerCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAccountIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (cc *CustomerCreate) AddBookingIDs(ids ...uuid.UUID) *CustomerCreate {
	cc.mutation.AddBookingIDs(ids...)
	return cc
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (cc *CustomerCreate) AddBookings(b ...*Booking) *CustomerCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddBookingIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Customer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CustomerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := customer.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := customer.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := customer.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Customer.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CitizenID(); !ok {
		return &ValidationError{Name: "citizen_id", err: errors.New(`ent: missing required field "Customer.citizen_id"`)}
	}
	if v, ok := cc.mutation.CitizenID(); ok {
		if err := customer.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Customer.citizen_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Customer.phone"`)}
	}
	if v, ok := cc.mutation.Phone(); ok {
		if err := customer.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Customer.phone": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Customer.address"`)}
	}
	if _, ok := cc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Customer.gender"`)}
	}
	if v, ok := cc.mutation.Gender(); ok {
		if err := customer.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Customer.gender": %w`, err)}
		}
	}
	if _, ok := cc.mutation.DateOfBirth(); !ok {
		return &ValidationError{Name: "date_of_birth", err: errors.New(`ent: missing required field "Customer.date_of_birth"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Customer.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Customer.updated_at"`)}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: customer.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.CitizenID(); ok {
		_spec.SetField(customer.FieldCitizenID, field.TypeString, value)
		_node.CitizenID = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.SetField(customer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := cc.mutation.Gender(); ok {
		_spec.SetField(customer.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := cc.mutation.DateOfBirth(); ok {
		_spec.SetField(customer.FieldDateOfBirth, field.TypeTime, value)
		_node.DateOfBirth = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AccountsTable,
			Columns: []string{customer.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BookingsTable,
			Columns: []string{customer.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
