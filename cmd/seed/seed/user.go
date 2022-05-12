package seed

import (
	"context"
	"log"
	"planetscale-witn-ent/ent"
)

// User generates  user data.
func User(ctx context.Context, client *ent.Client) {
	_, err := client.User.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TestUser failed to delete data: %v", err)
	}

	inputs := []ent.CreateUserInput{
		{
			Age:  20,
			Name: "Bob",
		},
		{
			Age:  30,
			Name: "Tom",
		},
		{
			Age:  23,
			Name: "Mary",
		},
	}
	bulk := make([]*ent.UserCreate, len(inputs))
	for i, t := range inputs {
		bulk[i] = client.User.Create().SetInput(t)
	}
	if _, err = client.User.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to seed data: %v", err)
	}
}
