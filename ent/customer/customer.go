// Code generated by ent, DO NOT EDIT.

package customer

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCitizenID holds the string denoting the citizen_id field in the database.
	FieldCitizenID = "citizen_id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldDateOfBirth holds the string denoting the date_of_birth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "accounts"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "customer_accounts"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "customer_bookings"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCitizenID,
	FieldPhone,
	FieldAddress,
	FieldGender,
	FieldDateOfBirth,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CitizenIDValidator is a validator for the "citizen_id" field. It is called by the builders before save.
	CitizenIDValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale, GenderOther:
		return nil
	default:
		return fmt.Errorf("customer: invalid enum value for gender field: %q", ge)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Gender(str)
	if err := GenderValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}
