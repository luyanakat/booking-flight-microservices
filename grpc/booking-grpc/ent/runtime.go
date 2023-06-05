// Code generated by ent, DO NOT EDIT.

package ent

import (
	"mock-project/grpc/booking-grpc/ent/booking"
	"mock-project/grpc/booking-grpc/ent/schema"
	"mock-project/grpc/booking-grpc/ent/ticket"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookingFields := schema.Booking{}.Fields()
	_ = bookingFields
	// bookingDescCode is the schema descriptor for code field.
	bookingDescCode := bookingFields[3].Descriptor()
	// booking.DefaultCode holds the default value on creation for the code field.
	booking.DefaultCode = bookingDescCode.Default.(func() string)
	// booking.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	booking.CodeValidator = bookingDescCode.Validators[0].(func(string) error)
	// bookingDescCreatedAt is the schema descriptor for created_at field.
	bookingDescCreatedAt := bookingFields[6].Descriptor()
	// booking.DefaultCreatedAt holds the default value on creation for the created_at field.
	booking.DefaultCreatedAt = bookingDescCreatedAt.Default.(func() time.Time)
	// bookingDescUpdatedAt is the schema descriptor for updated_at field.
	bookingDescUpdatedAt := bookingFields[7].Descriptor()
	// booking.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	booking.DefaultUpdatedAt = bookingDescUpdatedAt.Default.(func() time.Time)
	// booking.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	booking.UpdateDefaultUpdatedAt = bookingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bookingDescID is the schema descriptor for id field.
	bookingDescID := bookingFields[0].Descriptor()
	// booking.DefaultID holds the default value on creation for the id field.
	booking.DefaultID = bookingDescID.Default.(func() int64)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescCreatedAt is the schema descriptor for created_at field.
	ticketDescCreatedAt := ticketFields[2].Descriptor()
	// ticket.DefaultCreatedAt holds the default value on creation for the created_at field.
	ticket.DefaultCreatedAt = ticketDescCreatedAt.Default.(func() time.Time)
	// ticketDescUpdatedAt is the schema descriptor for updated_at field.
	ticketDescUpdatedAt := ticketFields[3].Descriptor()
	// ticket.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ticket.DefaultUpdatedAt = ticketDescUpdatedAt.Default.(func() time.Time)
	// ticket.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ticket.UpdateDefaultUpdatedAt = ticketDescUpdatedAt.UpdateDefault.(func() time.Time)
}