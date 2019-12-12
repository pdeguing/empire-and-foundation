package data

import (
	"context"
	"crypto/rand"
	dbsql "database/sql"
	"fmt"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

var Client *ent.Client
var DB *dbsql.DB

func open() (*ent.Client, error) {
	drv, err := sql.Open("postgres", "dbname=empire_and_foundation sslmode=disable")
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	DB = drv.DB()
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}

func init() {
	var err error
	Client, err = open()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return
}

// create a random UUID from RFC 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
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
