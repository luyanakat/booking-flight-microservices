package internal

import (
	"context"
	"mock-project/grpc/booking-grpc/ent"
)

func CreateTicketType(ctx context.Context, client *ent.Client) {
	// because too lazy to query from database and check data exist. I will set ID duplicate, so postgres will automatically not insert when run this after first time
	_ = client.Ticket.Create().SetID(1).SetName("First Class").Exec(ctx)
	_ = client.Ticket.Create().SetID(2).SetName("Economy Class").Exec(ctx)
}
