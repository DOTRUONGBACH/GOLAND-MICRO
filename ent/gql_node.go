// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"mock_project/ent/account"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"mock_project/ent/flight"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Account) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Account",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(a.Email); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Password); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "password",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Role); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "account.Role",
		Name:  "role",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Customer",
		Name: "acc_owner",
	}
	err = a.QueryAccOwner().
		Select(customer.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (b *Booking) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     b.ID,
		Type:   "Booking",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(b.TotalCost); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "float64",
		Name:  "total_cost",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.TotalTicket); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "total_ticket",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.Status); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "booking.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Flight",
		Name: "booking_flight",
	}
	err = b.QueryBookingFlight().
		Select(flight.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Customer",
		Name: "booking_owner",
	}
	err = b.QueryBookingOwner().
		Select(customer.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Customer) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Customer",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.CitizenID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "citizen_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Phone); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "phone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Address); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Gender); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "customer.Gender",
		Name:  "gender",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.DateOfBirth); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "date_of_birth",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Account",
		Name: "accounts",
	}
	err = c.QueryAccounts().
		Select(account.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Booking",
		Name: "bookings",
	}
	err = c.QueryBookings().
		Select(booking.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (f *Flight) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Flight",
		Fields: make([]*Field, 11),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(f.FlightCode); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "flight_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.From); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "from",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.To); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "to",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.DepartureDate); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "departure_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.DepartureTime); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "departure_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Duration); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "duration",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Capacity); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "capacity",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.AvailableSeat); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "int",
		Name:  "available_seat",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Status); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "flight.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Booking",
		Name: "belongs_to",
	}
	err = f.QueryBelongsTo().
		Select(booking.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case account.Table:
		query := c.Account.Query().
			Where(account.ID(id))
		query, err := query.CollectFields(ctx, "Account")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case booking.Table:
		query := c.Booking.Query().
			Where(booking.ID(id))
		query, err := query.CollectFields(ctx, "Booking")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case customer.Table:
		query := c.Customer.Query().
			Where(customer.ID(id))
		query, err := query.CollectFields(ctx, "Customer")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case flight.Table:
		query := c.Flight.Query().
			Where(flight.ID(id))
		query, err := query.CollectFields(ctx, "Flight")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case account.Table:
		query := c.Account.Query().
			Where(account.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Account")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case booking.Table:
		query := c.Booking.Query().
			Where(booking.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Booking")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case customer.Table:
		query := c.Customer.Query().
			Where(customer.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Customer")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case flight.Table:
		query := c.Flight.Query().
			Where(flight.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Flight")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
