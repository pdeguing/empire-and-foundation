package data

import (
	"fmt"
	"time"
)

type Planet struct {
	Id              int
	Uuid            string
	MetalStock      int64
	MetalMine       int
	UserId          int
	CreatedAt       time.Time
	LastMetalUpdate time.Time
	EndUpgradeTime  string
}

// GetMetalStock returns the current stock
func (planet *Planet) GetMetalStock() int64 {
	duration := int64(time.Since(planet.LastMetalUpdate) / time.Second)
	planet.MetalStock = planet.MetalStock + duration*planet.GetMetalRate()
	planet.LastMetalUpdate = time.Now().UTC()
	return planet.MetalStock
}

func (planet *Planet) UpdateMetalStock() {
	planet.GetMetalStock()

	statement := "update planets set metal_stock = $2, last_metal_update = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(planet.Id, planet.MetalStock, planet.LastMetalUpdate)
	return
}

// Computes the metal production rate, could potentially depend on multiple factors
func (planet *Planet) GetMetalRate() int64 {
	metalMine := int64(planet.MetalMine)
	rate := metalMine * 1
	return rate
}

func (planet *Planet) GetUpgradeTime() time.Duration {
	upgradeTime := time.Duration(planet.MetalMine) * time.Second
	return upgradeTime
}

func (planet *Planet) UpgradeMine() {
	time.AfterFunc(planet.GetUpgradeTime(), planet.ApplyUpgrade)
	planet.EndUpgradeTime = time.Now().Add(planet.GetUpgradeTime()).Format("Mon Jan 02 2006 15:04:05 GMT-0700")
	fmt.Println("timer started")
	fmt.Println(planet.EndUpgradeTime)
}

func (planet *Planet) ApplyUpgrade() {
	planet.EndUpgradeTime = time.Time{}.Format("Mon Jan 02 2006 15:04:05 GMT-0700")
	fmt.Println("timer done")
	fmt.Println(planet.EndUpgradeTime)
	// update metal stock
	planet.UpdateMetalStock()
	// change metal level
	planet.UpMetalMine()
}

// format the CreatedAt date to display nicely on the screen
func (planet *Planet) CreatedAtDate() string {
	return planet.CreatedAt.Format("Janv 2, 2006 at 3:04pm")
}

// Create a new planet
func (user *User) CreatePlanet() (planet Planet, err error) {
	statement := "insert into planets (uuid, metal_stock, metal_mine, user_id, created_at, last_metal_update) values ($1, $2, $3, $4, $5, $6) returning id, uuid, metal_stock, metal_mine, user_id, created_at, last_metal_update"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), "0", "0", user.Id, time.Now(), time.Now()).Scan(&planet.Id, &planet.Uuid, &planet.MetalStock, &planet.MetalMine, &planet.UserId, &planet.CreatedAt, &planet.LastMetalUpdate)
	return
}

// Get all planets in the database
func Planets() (planets []Planet, err error) {
	rows, err := Db.Query("SELECT id, uuid, metal_mine, metal_stock, user_id, created_at, last_metal_update FROM planets ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		p := Planet{}
		if err = rows.Scan(&p.Id, &p.Uuid, &p.MetalStock, &p.MetalMine, &p.UserId, &p.CreatedAt, &p.LastMetalUpdate); err != nil {
			return
		}
		planets = append(planets, p)
	}
	rows.Close()
	return
}

// Get a user's planet
func PlanetByUserId(userId int) (planet Planet, err error) {
	planet = Planet{}
	err = Db.QueryRow("SELECT id, uuid, metal_stock, metal_mine, user_id, created_at, last_metal_update FROM planets WHERE user_id = $1", userId).Scan(&planet.Id, &planet.Uuid, &planet.MetalStock, &planet.MetalMine, &planet.UserId, &planet.CreatedAt, &planet.LastMetalUpdate)
	if err != nil {
		return
	}
	planet.GetMetalStock()
	return
}

// Get a planet by UUID
func PlanetByUUID(uuid string) (planet Planet, err error) {
	planet = Planet{}
	err = Db.QueryRow("SELECT id, uuid, metal_stock, metal_mine, user_id, created_at, last_metal_update FROM planets WHERE uuid = $1", uuid).Scan(&planet.Id, &planet.Uuid, &planet.MetalStock, &planet.MetalMine, &planet.UserId, &planet.CreatedAt, &planet.LastMetalUpdate)
	planet.GetMetalStock()
	return
}

// Get the user who colonized this planet
func (planet *Planet) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", planet.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (planet *Planet) UpMetalMine() (err error) {
	statement := "update planets set metal_mine = $2 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	newLevel := planet.MetalMine + 1

	_, err = stmt.Exec(planet.Id, newLevel)
	return
}
