package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.Float("total_cost").Positive().Annotations(entgql.OrderField("TOTAL_COST"), entproto.Field(2)),
		field.Int("total_ticket").Positive().Default(1).Annotations(entgql.OrderField("TOTAL_TICKET"), entproto.Field(3)),
		field.Enum("status").Values("Success", "Fail", "Canceled").Annotations(entgql.OrderField("STATUS"), entproto.Field(4)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(5)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(6)),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking_flight", Flight.Type).
			 Ref("belongs_to").Unique(),
		edge.From("booking_owner", Customer.Type).
			 Ref("bookings").Unique(),
	}
}

func (Booking) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
