package data

import (
	"log"
	"math/rand"
	"context"
	"io/ioutil"
	"strings"
	"os"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/goombaio/namegenerator"
)

func randomPlanetType(r *rand.Rand) planet.PlanetType {
	n := r.Intn(9)

	if n < 1 {
		return planet.PlanetTypeHabitable
	} else if n < 3 {
		return planet.PlanetTypeIceGiant
	} else if n < 6 {
		return planet.PlanetTypeGasGiant
	} else {
		return planet.PlanetTypeMineral
	}
}

func getPositionCode(r int, s int, o int, su int) int {
	return su + o << 4 + s << 8 + r << 12

}

func randomPlanetName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name
}

func randomPlanetSkin(r *rand.Rand, planetSkins []os.FileInfo) string {
	n := r.Intn(len(planetSkins) - 1)

	planetSkin := planetSkins[n].Name()
	planetSkin = strings.TrimSuffix(planetSkin, ".png")

	return planetSkin
}

func generateEntity(region int, system int, orbit int, suborbit int, planetType planet.PlanetType, planetSkin string) {
	positionCode := getPositionCode(region, system, orbit, suborbit)
	planetName := randomPlanetName()

	log.Println("create planet:", planetName, planetType, positionCode, planetSkin, "(with:", region, system, orbit, suborbit, ")")

	_ = Client.Planet.
		Create().
		SetName(planetName).
		SetPlanetType(planetType).
		SetRegionCode(region).
		SetSystemCode(system).
		SetOrbitCode(orbit).
		SetSuborbitCode(suborbit).
		SetPositionCode(positionCode).
		SetPlanetSkin(planetSkin).
		SaveX(context.Background())

}

func generatePlanet(r *rand.Rand, planetSkins []os.FileInfo, region int, system int, orbit int, suborbit int) bool {
	n := r.Intn(9)

	if n < 4 {
		planetType := randomPlanetType(r)
		planetSkin := randomPlanetSkin(r, planetSkins)
		generateEntity(region, system, orbit, suborbit, planetType, planetSkin)
		return true
	}
	return false
}

func generateRegion(region int) {
	r := rand.New(rand.NewSource(42))

	planetSkins, err := ioutil.ReadDir("public/images/planet-dashboards")
	if err != nil {
		log.Fatal(err)
	}

	generated := 0

	for system := 0; system < 256; system++ {
		for orbit := 1; orbit < 16; orbit++ {
			if generatePlanet(r, planetSkins, region, system, orbit, 0) {
				generated++
//				for suborbit := 1; suborbit < 16; suborbit++ {
//					generatePlanet(r, planetSkins, region, system, orbit, suborbit)
//				}
			}
		}
	}

	log.Println(generated, "planets generated")
}
