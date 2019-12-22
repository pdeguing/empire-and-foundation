package data

import (
	"log"
	"math/rand"
	"context"
	"io/ioutil"
	"strings"

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

func randomPlanetSkin(r *rand.Rand) string {
	files, err := ioutil.ReadDir("public/images/planet-dashboards")
	if err != nil {
		log.Fatal(err)
	}

	n := r.Intn(len(files) - 1)

	planetSkin := files[n].Name()
	planetSkin = strings.TrimSuffix(planetSkin, ".png")

	return planetSkin
}

func generateEntity(region int, system int, orbit int, suborbit int, planetType planet.PlanetType, planetSkin string) {
	positionCode := getPositionCode(region, system, orbit, suborbit)

	log.Println("create planet:", planetType, positionCode, planetSkin, "(with:", region, system, orbit, suborbit, ")")

	_ = Client.Planet.
		Create().
		SetPlanetType(planetType).
		SetRegionCode(region).
		SetSystemCode(system).
		SetOrbitCode(orbit).
		SetSuborbitCode(suborbit).
		SetPositionCode(positionCode).
		SetPlanetSkin(planetSkin).
		SaveX(context.Background())

}

func generatePlanet(r *rand.Rand, region int, system int, orbit int, suborbit int) bool {
	n := r.Intn(9)

	if n < 2 {
		planetType := randomPlanetType(r)
		planetSkin := randomPlanetSkin(r)
		generateEntity(region, system, orbit, suborbit, planetType, planetSkin)
		return true
	}
	return false
}

func generateRegion(region int) {
	r := rand.New(rand.NewSource(42))

	for system := 0; system < 256; system++ {
		for orbit := 1; orbit < 16; orbit++ {
			if generatePlanet(r, region, system, orbit, 0) {
				for suborbit := 1; suborbit < 16; suborbit++ {
					generatePlanet(r, region, system, orbit, suborbit)
				}
			}
		}
	}
}
