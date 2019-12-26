package data

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"testing"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/schema"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

var testDatabaseFileCounter int

// WithTestDatabase sets the global Db and Client to a temporary in-memory database.
func WithTestDatabase() {
	DB.Close()

	// TODO: Does this leak memory or does the DB.Close() take care of it?
	testDatabaseFileCounter++
	drv, err := sql.Open("sqlite3", "file:file_"+strconv.Itoa(testDatabaseFileCounter)+".db?cache=shared&mode=memory&_foreign_keys=1")
	if err != nil {
		panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	DB = drv.DB()
	Client = ent.NewClient(ent.Driver(drv))

	err = Migrate(context.Background(), Client)
	if err != nil {
		panic(err)
	}
}

func newTestPlanet(ctx context.Context, client *ent.Client) *ent.Planet {
	p, err := client.Planet.Create().
		SetName("Foobar planet").
		SetMetal(100000).
		SetSilica(10000).
		SetHydrogen(10000).
		Save(ctx)
	if err != nil {
		panic(err)
	}
	return p
}

func entityEquals(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	// Convert the entities to json so the testing library can do a
	// line-by-line diff which only includes public properties.
	jsonExpected, err := json.MarshalIndent(expected, "", "    ")
	if err != nil {
		panic(err)
	}
	jsonActual, err := json.MarshalIndent(actual, "", "    ")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, string(jsonExpected), string(jsonActual), msgAndArgs...)
}

func TestActions_ContainsAllSchemaDefinedActions(t *testing.T) {
	var actns []string
	for _, field := range (schema.Timer{}).Fields() {
		d := field.Descriptor()
		if d.Name == "action" {
			actns = d.Enums
			break
		}
	}
	for _, a := range actns {
		assert.Contains(t, actions, timer.Action(a), "Map \"actions\" does not contain the action %q", a)
	}
}

func TestActions_StartFunctionAltersPlanetConsistentlyWithDatabase(t *testing.T) {
	WithTestDatabase()
	ctx := context.Background()
	for name, a := range actions {
		WithTx(ctx, Client, func(tx *ent.Tx) error {
			pStruct := newTestPlanet(ctx, tx.Client())
			err := a.Start(ctx, tx, pStruct)
			assert.NoError(t, err)

			pDatabase, err := tx.Planet.
				Query().
				Where(planet.IDEQ(pStruct.ID)).
				Only(ctx)
			if err != nil {
				panic(err)
			}

			// Make the updated at field the same to ignore it in the comparison.
			pStruct.UpdatedAt = pDatabase.UpdatedAt

			entityEquals(t, pDatabase, pStruct, "Start function for %q doesn't make the same changes to the planet struct as to the planet stored in the database.", name)

			return errors.New("rollback")
		})
	}
}

func TestActions_CompleteFunctionAltersPlanetConsistentlyWithDatabase(t *testing.T) {
	WithTestDatabase()
	ctx := context.Background()
	for name, a := range actions {
		WithTx(ctx, Client, func(tx *ent.Tx) error {
			pStruct := newTestPlanet(ctx, tx.Client())
			err := a.Complete(ctx, tx, pStruct)
			assert.NoError(t, err)

			pDatabase, err := tx.Planet.
				Query().
				Where(planet.IDEQ(pStruct.ID)).
				Only(ctx)
			if err != nil {
				panic(err)
			}

			// Make the updated at field the same to ignore it in the comparison.
			pStruct.UpdatedAt = pDatabase.UpdatedAt

			entityEquals(t, pDatabase, pStruct, "Start function for %q doesn't make the same changes to the planet struct as to the planet stored in the database.", name)

			return errors.New("rollback")
		})
	}
}

func TestActions_CancelFunctionAltersPlanetConsistentlyWithDatabase(t *testing.T) {
	WithTestDatabase()
	ctx := context.Background()
	for name, a := range actions {
		WithTx(ctx, Client, func(tx *ent.Tx) error {
			pStruct := newTestPlanet(ctx, tx.Client())
			err := a.Cancel(ctx, tx, pStruct)
			assert.NoError(t, err)

			pDatabase, err := tx.Planet.
				Query().
				Where(planet.IDEQ(pStruct.ID)).
				Only(ctx)
			if err != nil {
				panic(err)
			}

			// Make the updated at field the same to ignore it in the comparison.
			pStruct.UpdatedAt = pDatabase.UpdatedAt

			entityEquals(t, pDatabase, pStruct, "Start function for %q doesn't make the same changes to the planet struct as to the planet stored in the database.", name)

			return errors.New("rollback")
		})
	}
}
