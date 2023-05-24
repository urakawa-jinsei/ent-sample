package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"

	_ "github.com/lib/pq"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=user dbname=db password=pass")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// CreateUserを呼び出す
	ctx := context.Background()
	user, err := CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	fmt.Println("User created:", user)
}
