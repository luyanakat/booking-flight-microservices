package internal

import (
	"context"
	"mock-project/grpc/user-grpc/ent"
	"mock-project/grpc/user-grpc/internal/auth"
)

func CreateRole(ctx context.Context, client *ent.Client) {
	// just for init data

	// because too lazy to query from database and check data exist. I will set ID duplicate, so postgres will automatically not insert when run this after first time
	_ = client.AccessLevel.Create().SetID(1).SetName("Guest").Exec(ctx)
	_ = client.AccessLevel.Create().SetID(2).SetName("User").Exec(ctx)
	_ = client.AccessLevel.Create().SetID(3).SetName("Admin").Exec(ctx)
	passHashed, _ := auth.HashPassword("admin@123")
	_ = client.User.Create().SetEmail("admin@techvify.com").SetPassword(passHashed).SetAccessID(3).SetCustomerID(0).Exec(ctx)
}
