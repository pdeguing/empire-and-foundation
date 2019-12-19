package data

import (
	"log"
	"math/rand"
	"context"

	"github.com/pdeguing/empire-and-foundation/ent/planet"
)

func randomPlanetType(r *rand.Rand) planet.PlanetType {
	n := r.Intn(9)

	if n < 2 {
		return planet.PlanetTypeHabitable
	} else if n > 6 {
		return planet.PlanetTypeGaseous
	} else {
		return planet.PlanetTypeMineral
	}
}

func getPositionCode(r int, s int, o int, su int) int {
	return su + o << 4 + s << 8 + r << 12

}

func generateEntity(region int, system int, orbit int, suborbit int, planetType planet.PlanetType) {
	positionCode := getPositionCode(region, system, orbit, suborbit)

	_ = Client.Planet.
		Create().
		SetPlanetType(planetType).
		SetRegionCode(region).
		SetSystemCode(system).
		SetOrbitCode(orbit).
		SetSuborbitCode(suborbit).
		SetPositionCode(positionCode).
		SaveX(context.Background())
	log.Println("created planet: %d, %d", planetType, positionCode)

}

func generatePlanet(r *rand.Rand, region int, system int, orbit int, suborbit int) {
	n := r.Intn(9)

	if n < 2 {
		planetType := randomPlanetType(r)
		generateEntity(region, system, orbit, suborbit, planetType)
	}
}

func generateRegion(region int) {
	r := rand.New(rand.NewSource(42))

	for system := 0; system < 256; system++ {
		for orbit := 1; orbit < 16; orbit++ {
			generatePlanet(r, region, system, orbit, 0)
			for suborbit := 0; suborbit < 16; suborbit++ {
				generatePlanet(r, region, system, orbit, suborbit)
			}
		}
	}
}
