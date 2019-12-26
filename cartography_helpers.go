package main

import (
	"net/http"
	"fmt"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
)

func regionPlanets(w http.ResponseWriter, r *http.Request) ([]*ent.Planet, error) {
	p, err := data.Client.Planet.Query().
		Order(ent.Asc(planet.FieldPositionCode)).
		All(r.Context())

	if _, ok := err.(*ent.ErrNotFound); ok {
		return nil, newNotFoundError(err)
	}

	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("could not retrieve user's planet from database: %v", err))
	}

	return p, err
}
