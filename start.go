package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"entdemo/ent"
	"entdemo/ent/car"
	"entdemo/ent/user"

	_ "github.com/lib/pq"
)

// ユーザーを作成する
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

// ユーザーを一人取り出す
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

// 車とユーザーを登録する
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

// 車を取り出す
func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func main() {
	ctx := context.Background()
	client, err := ent.Open("postgres", "host=localhost port=5432 user=user dbname=db password=pass")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// CreateUserを呼び出す
	// user, err := CreateUser(ctx, client)
	// if err != nil {
	// 	log.Fatalf("failed to create user: %v", err)
	// }
	// fmt.Println("User created:", user)

	// QueryUserを呼び出す
	user, err := QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed to query users: %v", err)
	}
	fmt.Println("User returned:", user)

	// CreateCarsを呼び出す
	// user, err := CreateCars(ctx, client)
	// if err != nil {
	// 	log.Fatalf("failed to create cars: %v", err)
	// }
	// fmt.Println("Cars created:", user)

	// QueryCarsを呼び出す
	err = QueryCars(ctx, user)
	if err != nil {
		log.Fatalf("failed to create cars: %v", err)
	}
}
