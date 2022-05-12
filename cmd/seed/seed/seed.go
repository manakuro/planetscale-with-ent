package seed

import (
	"context"
	"log"
	"planetscale-witn-ent/datastore"
	"planetscale-witn-ent/ent"
)

// Seed truncate table and generates new data.
func Seed() {
	client := newDBClient()
	defer client.Close()

	ctx := context.Background()

	client.DisableForeignKeyChecks()
	client.DisableSQLSafeUpdates()

	User(ctx, client)

	client.EnableForeignKeyChecks()
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient(datastore.NewClientOptions{})
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}
