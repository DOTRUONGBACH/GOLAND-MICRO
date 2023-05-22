// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"mock_project/ent/flight"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetTotalCost sets the "total_cost" field.
func (bc *BookingCreate) SetTotalCost(f float64) *BookingCreate {
	bc.mutation.SetTotalCost(f)
	return bc
}

// SetTotalTicket sets the "total_ticket" field.
func (bc *BookingCreate) SetTotalTicket(i int) *BookingCreate {
	bc.mutation.SetTotalTicket(i)
	return bc
}

// SetNillableTotalTicket sets the "total_ticket" field if the given value is not nil.
func (bc *BookingCreate) SetNillableTotalTicket(i *int) *BookingCreate {
	if i != nil {
		bc.SetTotalTicket(*i)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BookingCreate) SetStatus(b booking.Status) *BookingCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookingCreate) SetCreatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableCreatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BookingCreate) SetUpdatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableUpdatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BookingCreate) SetID(u uuid.UUID) *BookingCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BookingCreate) SetNillableID(u *uuid.UUID) *BookingCreate {
	if u != nil {
		bc.SetID(*u)
	}
	return bc
}

// SetBookingFlightID sets the "booking_flight" edge to the Flight entity by ID.
func (bc *BookingCreate) SetBookingFlightID(id uuid.UUID) *BookingCreate {
	bc.mutation.SetBookingFlightID(id)
	return bc
}

// SetNillableBookingFlightID sets the "booking_flight" edge to the Flight entity by ID if the given value is not nil.
func (bc *BookingCreate) SetNillableBookingFlightID(id *uuid.UUID) *BookingCreate {
	if id != nil {
		bc = bc.SetBookingFlightID(*id)
	}
	return bc
}

// SetBookingFlight sets the "booking_flight" edge to the Flight entity.
func (bc *BookingCreate) SetBookingFlight(f *Flight) *BookingCreate {
	return bc.SetBookingFlightID(f.ID)
}

// SetBookingOwnerID sets the "booking_owner" edge to the Customer entity by ID.
func (bc *BookingCreate) SetBookingOwnerID(id uuid.UUID) *BookingCreate {
	bc.mutation.SetBookingOwnerID(id)
	return bc
}

// SetNillableBookingOwnerID sets the "booking_owner" edge to the Customer entity by ID if the given value is not nil.
func (bc *BookingCreate) SetNillableBookingOwnerID(id *uuid.UUID) *BookingCreate {
	if id != nil {
		bc = bc.SetBookingOwnerID(*id)
	}
	return bc
}

// SetBookingOwner sets the "booking_owner" edge to the Customer entity.
func (bc *BookingCreate) SetBookingOwner(c *Customer) *BookingCreate {
	return bc.SetBookingOwnerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	var (
		err  error
		node *Booking
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Booking)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BookingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookingCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookingCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookingCreate) defaults() {
	if _, ok := bc.mutation.TotalTicket(); !ok {
		v := booking.DefaultTotalTicket
		bc.mutation.SetTotalTicket(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := booking.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := booking.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := booking.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookingCreate) check() error {
	if _, ok := bc.mutation.TotalCost(); !ok {
		return &ValidationError{Name: "total_cost", err: errors.New(`ent: missing required field "Booking.total_cost"`)}
	}
	if v, ok := bc.mutation.TotalCost(); ok {
		if err := booking.TotalCostValidator(v); err != nil {
			return &ValidationError{Name: "total_cost", err: fmt.Errorf(`ent: validator failed for field "Booking.total_cost": %w`, err)}
		}
	}
	if _, ok := bc.mutation.TotalTicket(); !ok {
		return &ValidationError{Name: "total_ticket", err: errors.New(`ent: missing required field "Booking.total_ticket"`)}
	}
	if v, ok := bc.mutation.TotalTicket(); ok {
		if err := booking.TotalTicketValidator(v); err != nil {
			return &ValidationError{Name: "total_ticket", err: fmt.Errorf(`ent: validator failed for field "Booking.total_ticket": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Booking.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Booking.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Booking.updated_at"`)}
	}
	return nil
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		_node = &Booking{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: booking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: booking.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.TotalCost(); ok {
		_spec.SetField(booking.FieldTotalCost, field.TypeFloat64, value)
		_node.TotalCost = value
	}
	if value, ok := bc.mutation.TotalTicket(); ok {
		_spec.SetField(booking.FieldTotalTicket, field.TypeInt, value)
		_node.TotalTicket = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.BookingFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingFlightTable,
			Columns: []string{booking.BookingFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flight.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.flight_belongs_to = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BookingOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingOwnerTable,
			Columns: []string{booking.BookingOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_bookings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookingCreateBulk is the builder for creating many Booking entities in bulk.
type BookingCreateBulk struct {
	config
	builders []*BookingCreate
}

// Save creates the Booking entities in the database.
func (bcb *BookingCreateBulk) Save(ctx context.Context) ([]*Booking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Booking, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookingMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookingCreateBulk) SaveX(ctx context.Context) []*Booking {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookingCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookingCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
