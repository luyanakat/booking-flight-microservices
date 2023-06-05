package dbrepo

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mock-project/grpc/user-grpc/ent"
	internal "mock-project/grpc/user-grpc/internal/init"
	repository "mock-project/grpc/user-grpc/repo"

	_ "github.com/lib/pq"
)

type postgresDBRepo struct {
	client *ent.Client
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.database"),
	)
}

func NewPostgresRepo(ctx context.Context) (repository.UserRepository, error) {
	client := Open(ConnectionString())

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// init role data
	internal.CreateRole(ctx, client)

	return &postgresDBRepo{client: client}, nil
}
