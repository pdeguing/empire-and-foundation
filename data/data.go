package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	entsql "github.com/facebookincubator/ent/dialect/sql"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/migrate"
	"github.com/pkg/errors"
)

// Client is an ent ORM client with which database queries can be made.
var Client *ent.Client

// DB is the raw SQL DB handle to the database. It is recommended to use
// Client, but in some cases the raw handle is needed. For instance, in
// use of external libraries.
var DB *sql.DB

// InitDatabaseConnection makes a connection to the database and
// sets Client and DB so they can be used by the rest of the application.
func InitDatabaseConnection(driver, source string, debug bool) error {
	drv, err := entsql.Open(driver, source)
	if err != nil {
		return fmt.Errorf("failed opening database connection: %w", err)
	}

	// Get the underlying sql.DB object of the driver.
	DB = drv.DB()
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)

	Client = ent.NewClient(ent.Driver(drv))
	if debug {
		Client = Client.Debug()
	}
	return nil
}

// Migrate runs the database schema migrations.
func Migrate(ctx context.Context, client *ent.Client, dropIndex, dropColumn bool) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(dropIndex),
		migrate.WithDropColumn(dropColumn),
	)
	if err != nil {
		return fmt.Errorf("unable to migrate database: %w", err)
	}
	return nil
}

// WithTx wraps fn in a transaction that automatically rolls back
// on error or when a panic occurs.
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
