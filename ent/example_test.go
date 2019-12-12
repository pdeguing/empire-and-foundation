// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExamplePlanet() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the planet's edges.

	// create planet vertex with its edges.
	pl := client.Planet.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetMetalStock(1).
		SetMetalMine(1).
		SetLastMetalUpdate(time.Now()).
		SaveX(ctx)
	log.Println("planet created:", pl)

	// query edges.

	// Output:
}
func ExampleSession() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the session's edges.

	// create session vertex with its edges.
	s := client.Session.
		Create().
		SetToken("string").
		SetData(nil).
		SetExpiry(time.Now()).
		SaveX(ctx)
	log.Println("session created:", s)

	// query edges.

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	pl0 := client.Planet.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetMetalStock(1).
		SetMetalMine(1).
		SetLastMetalUpdate(time.Now()).
		SaveX(ctx)
	log.Println("planet created:", pl0)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetUsername("string").
		SetEmail("string").
		SetPassword("string").
		AddPlanets(pl0).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	pl0, err = u.QueryPlanets().First(ctx)
	if err != nil {
		log.Fatalf("failed querying planets: %v", err)
	}
	log.Println("planets found:", pl0)

	// Output:
}
