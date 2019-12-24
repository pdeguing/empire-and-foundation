package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
)

func regionPlanets(w http.ResponseWriter, r *http.Request) ([]*ent.Planet, bool) {
	p, err := data.Client.Planet.Query().
		Order(ent.Asc(planet.FieldPositionCode)).
		All(r.Context())

	if _, ok := err.(*ent.ErrNotFound); ok {
		serveNotFoundError(w, r)
		return nil, false
	}

	if err != nil {
		serveInternalServerError(w, r, err, "Could not retrieve user's planet from database")
		return nil, false
	}

	return p, true
}
