package dbrepo

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mock-project/grpc/flight-grpc/ent"
	repository "mock-project/grpc/flight-grpc/repo"

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

func NewPostgresRepo() (repository.FlightRepository, error) {
	client := Open(ConnectionString())
	log.Println(ConnectionString())

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &postgresDBRepo{client: client}, nil
}
