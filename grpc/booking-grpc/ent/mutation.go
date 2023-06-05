// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock-project/grpc/booking-grpc/ent/booking"
	"mock-project/grpc/booking-grpc/ent/predicate"
	"mock-project/grpc/booking-grpc/ent/ticket"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBooking = "Booking"
	TypeTicket  = "Ticket"
)

// BookingMutation represents an operation that mutates the Booking nodes in the graph.
type BookingMutation struct {
	config
	op             Op
	typ            string
	id             *int64
	customer_id    *int64
	addcustomer_id *int64
	flight_id      *int64
	addflight_id   *int64
	code           *string
	status         *booking.Status
	ticket_id      *int64
	addticket_id   *int64
	created_at     *time.Time
	updated_at     *time.Time
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Booking, error)
	predicates     []predicate.Booking
}

var _ ent.Mutation = (*BookingMutation)(nil)

// bookingOption allows management of the mutation configuration using functional options.
type bookingOption func(*BookingMutation)

// newBookingMutation creates new mutation for the Booking entity.
func newBookingMutation(c config, op Op, opts ...bookingOption) *BookingMutation {
	m := &BookingMutation{
		config:        c,
		op:            op,
		typ:           TypeBooking,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBookingID sets the ID field of the mutation.
func withBookingID(id int64) bookingOption {
	return func(m *BookingMutation) {
		var (
			err   error
			once  sync.Once
			value *Booking
		)
		m.oldValue = func(ctx context.Context) (*Booking, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Booking.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBooking sets the old Booking of the mutation.
func withBooking(node *Booking) bookingOption {
	return func(m *BookingMutation) {
		m.oldValue = func(context.Context) (*Booking, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BookingMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BookingMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Booking entities.
func (m *BookingMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BookingMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BookingMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Booking.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCustomerID sets the "customer_id" field.
func (m *BookingMutation) SetCustomerID(i int64) {
	m.customer_id = &i
	m.addcustomer_id = nil
}

// CustomerID returns the value of the "customer_id" field in the mutation.
func (m *BookingMutation) CustomerID() (r int64, exists bool) {
	v := m.customer_id
	if v == nil {
		return
	}
	return *v, true
}

// OldCustomerID returns the old "customer_id" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldCustomerID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCustomerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCustomerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustomerID: %w", err)
	}
	return oldValue.CustomerID, nil
}

// AddCustomerID adds i to the "customer_id" field.
func (m *BookingMutation) AddCustomerID(i int64) {
	if m.addcustomer_id != nil {
		*m.addcustomer_id += i
	} else {
		m.addcustomer_id = &i
	}
}

// AddedCustomerID returns the value that was added to the "customer_id" field in this mutation.
func (m *BookingMutation) AddedCustomerID() (r int64, exists bool) {
	v := m.addcustomer_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetCustomerID resets all changes to the "customer_id" field.
func (m *BookingMutation) ResetCustomerID() {
	m.customer_id = nil
	m.addcustomer_id = nil
}

// SetFlightID sets the "flight_id" field.
func (m *BookingMutation) SetFlightID(i int64) {
	m.flight_id = &i
	m.addflight_id = nil
}

// FlightID returns the value of the "flight_id" field in the mutation.
func (m *BookingMutation) FlightID() (r int64, exists bool) {
	v := m.flight_id
	if v == nil {
		return
	}
	return *v, true
}

// OldFlightID returns the old "flight_id" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldFlightID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFlightID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFlightID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFlightID: %w", err)
	}
	return oldValue.FlightID, nil
}

// AddFlightID adds i to the "flight_id" field.
func (m *BookingMutation) AddFlightID(i int64) {
	if m.addflight_id != nil {
		*m.addflight_id += i
	} else {
		m.addflight_id = &i
	}
}

// AddedFlightID returns the value that was added to the "flight_id" field in this mutation.
func (m *BookingMutation) AddedFlightID() (r int64, exists bool) {
	v := m.addflight_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetFlightID resets all changes to the "flight_id" field.
func (m *BookingMutation) ResetFlightID() {
	m.flight_id = nil
	m.addflight_id = nil
}

// SetCode sets the "code" field.
func (m *BookingMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *BookingMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *BookingMutation) ResetCode() {
	m.code = nil
}

// SetStatus sets the "status" field.
func (m *BookingMutation) SetStatus(b booking.Status) {
	m.status = &b
}

// Status returns the value of the "status" field in the mutation.
func (m *BookingMutation) Status() (r booking.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldStatus(ctx context.Context) (v booking.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *BookingMutation) ResetStatus() {
	m.status = nil
}

// SetTicketID sets the "ticket_id" field.
func (m *BookingMutation) SetTicketID(i int64) {
	m.ticket_id = &i
	m.addticket_id = nil
}

// TicketID returns the value of the "ticket_id" field in the mutation.
func (m *BookingMutation) TicketID() (r int64, exists bool) {
	v := m.ticket_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTicketID returns the old "ticket_id" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldTicketID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTicketID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTicketID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTicketID: %w", err)
	}
	return oldValue.TicketID, nil
}

// AddTicketID adds i to the "ticket_id" field.
func (m *BookingMutation) AddTicketID(i int64) {
	if m.addticket_id != nil {
		*m.addticket_id += i
	} else {
		m.addticket_id = &i
	}
}

// AddedTicketID returns the value that was added to the "ticket_id" field in this mutation.
func (m *BookingMutation) AddedTicketID() (r int64, exists bool) {
	v := m.addticket_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetTicketID resets all changes to the "ticket_id" field.
func (m *BookingMutation) ResetTicketID() {
	m.ticket_id = nil
	m.addticket_id = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *BookingMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *BookingMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *BookingMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *BookingMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *BookingMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Booking entity.
// If the Booking object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookingMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *BookingMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the BookingMutation builder.
func (m *BookingMutation) Where(ps ...predicate.Booking) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BookingMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BookingMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Booking, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BookingMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BookingMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Booking).
func (m *BookingMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BookingMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.customer_id != nil {
		fields = append(fields, booking.FieldCustomerID)
	}
	if m.flight_id != nil {
		fields = append(fields, booking.FieldFlightID)
	}
	if m.code != nil {
		fields = append(fields, booking.FieldCode)
	}
	if m.status != nil {
		fields = append(fields, booking.FieldStatus)
	}
	if m.ticket_id != nil {
		fields = append(fields, booking.FieldTicketID)
	}
	if m.created_at != nil {
		fields = append(fields, booking.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, booking.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BookingMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case booking.FieldCustomerID:
		return m.CustomerID()
	case booking.FieldFlightID:
		return m.FlightID()
	case booking.FieldCode:
		return m.Code()
	case booking.FieldStatus:
		return m.Status()
	case booking.FieldTicketID:
		return m.TicketID()
	case booking.FieldCreatedAt:
		return m.CreatedAt()
	case booking.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BookingMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case booking.FieldCustomerID:
		return m.OldCustomerID(ctx)
	case booking.FieldFlightID:
		return m.OldFlightID(ctx)
	case booking.FieldCode:
		return m.OldCode(ctx)
	case booking.FieldStatus:
		return m.OldStatus(ctx)
	case booking.FieldTicketID:
		return m.OldTicketID(ctx)
	case booking.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case booking.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Booking field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookingMutation) SetField(name string, value ent.Value) error {
	switch name {
	case booking.FieldCustomerID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustomerID(v)
		return nil
	case booking.FieldFlightID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFlightID(v)
		return nil
	case booking.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case booking.FieldStatus:
		v, ok := value.(booking.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case booking.FieldTicketID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTicketID(v)
		return nil
	case booking.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case booking.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Booking field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BookingMutation) AddedFields() []string {
	var fields []string
	if m.addcustomer_id != nil {
		fields = append(fields, booking.FieldCustomerID)
	}
	if m.addflight_id != nil {
		fields = append(fields, booking.FieldFlightID)
	}
	if m.addticket_id != nil {
		fields = append(fields, booking.FieldTicketID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BookingMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case booking.FieldCustomerID:
		return m.AddedCustomerID()
	case booking.FieldFlightID:
		return m.AddedFlightID()
	case booking.FieldTicketID:
		return m.AddedTicketID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookingMutation) AddField(name string, value ent.Value) error {
	switch name {
	case booking.FieldCustomerID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCustomerID(v)
		return nil
	case booking.FieldFlightID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFlightID(v)
		return nil
	case booking.FieldTicketID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTicketID(v)
		return nil
	}
	return fmt.Errorf("unknown Booking numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BookingMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BookingMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BookingMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Booking nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BookingMutation) ResetField(name string) error {
	switch name {
	case booking.FieldCustomerID:
		m.ResetCustomerID()
		return nil
	case booking.FieldFlightID:
		m.ResetFlightID()
		return nil
	case booking.FieldCode:
		m.ResetCode()
		return nil
	case booking.FieldStatus:
		m.ResetStatus()
		return nil
	case booking.FieldTicketID:
		m.ResetTicketID()
		return nil
	case booking.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case booking.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Booking field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BookingMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BookingMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BookingMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BookingMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BookingMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BookingMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BookingMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Booking unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BookingMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Booking edge %s", name)
}

// TicketMutation represents an operation that mutates the Ticket nodes in the graph.
type TicketMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Ticket, error)
	predicates    []predicate.Ticket
}

var _ ent.Mutation = (*TicketMutation)(nil)

// ticketOption allows management of the mutation configuration using functional options.
type ticketOption func(*TicketMutation)

// newTicketMutation creates new mutation for the Ticket entity.
func newTicketMutation(c config, op Op, opts ...ticketOption) *TicketMutation {
	m := &TicketMutation{
		config:        c,
		op:            op,
		typ:           TypeTicket,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTicketID sets the ID field of the mutation.
func withTicketID(id int64) ticketOption {
	return func(m *TicketMutation) {
		var (
			err   error
			once  sync.Once
			value *Ticket
		)
		m.oldValue = func(ctx context.Context) (*Ticket, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Ticket.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTicket sets the old Ticket of the mutation.
func withTicket(node *Ticket) ticketOption {
	return func(m *TicketMutation) {
		m.oldValue = func(context.Context) (*Ticket, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TicketMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TicketMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Ticket entities.
func (m *TicketMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TicketMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TicketMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Ticket.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *TicketMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TicketMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Ticket entity.
// If the Ticket object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TicketMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TicketMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TicketMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TicketMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Ticket entity.
// If the Ticket object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TicketMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TicketMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TicketMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TicketMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Ticket entity.
// If the Ticket object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TicketMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TicketMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the TicketMutation builder.
func (m *TicketMutation) Where(ps ...predicate.Ticket) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TicketMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TicketMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Ticket, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TicketMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TicketMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Ticket).
func (m *TicketMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TicketMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, ticket.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, ticket.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, ticket.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TicketMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case ticket.FieldName:
		return m.Name()
	case ticket.FieldCreatedAt:
		return m.CreatedAt()
	case ticket.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TicketMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case ticket.FieldName:
		return m.OldName(ctx)
	case ticket.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case ticket.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Ticket field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TicketMutation) SetField(name string, value ent.Value) error {
	switch name {
	case ticket.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case ticket.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case ticket.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Ticket field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TicketMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TicketMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TicketMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Ticket numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TicketMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TicketMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TicketMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Ticket nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TicketMutation) ResetField(name string) error {
	switch name {
	case ticket.FieldName:
		m.ResetName()
		return nil
	case ticket.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case ticket.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Ticket field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TicketMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TicketMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TicketMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TicketMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TicketMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TicketMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TicketMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Ticket unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TicketMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Ticket edge %s", name)
}