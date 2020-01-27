package data

import (
	"context"
	"github.com/pdeguing/empire-and-foundation/ent"
	"golang.org/x/crypto/bcrypt"
	"syreclabs.com/go/faker"
	"time"
)

type userFactory struct {
	ent.User
	client *ent.Client
	ctx    context.Context
}

// NewUserFactory creates a factory initialized with random data.
func NewUserFactory() *userFactory {
	f := userFactory{}
	f.Email = faker.Internet().SafeEmail()
	f.Username = faker.Internet().UserName()
	f.Password = fakePassword()
	f.CreatedAt = faker.Time().Backward(time.Hour * 24 * 365)
	f.UpdatedAt = faker.Time().Backward(time.Since(f.CreatedAt))
	f.client = Client
	f.ctx = context.Background()
	return &f
}

// WithPlanet adds a planet to the user.
func (f *userFactory) WithPlanet() *userFactory {
	p := NewPlanetFactory().Client(f.client).MustCreate()
	f.Edges.Planets = append(f.Edges.Planets, p)
	return f
}

// WithPlanets adds n planets to the user.
func (f *userFactory) WithPlanets(n uint) *userFactory {
	for i := uint(0); i < n; i++ {
		f.WithPlanet()
	}
	return f
}

// Client uses c to create the entity.
func (f *userFactory) Client(c *ent.Client) *userFactory {
	f.client = c
	return f
}

// Tx uses c to create the entity.
func (f *userFactory) Tx(tx *ent.Tx) *userFactory {
	return f.Client(tx.Client())
}

// InContext executes the Create method in ctx.
func (f *userFactory) InContext(ctx context.Context) *userFactory {
	f.ctx = ctx
	return f
}

// Create returns the user struct, which is saved to the database.
func (f *userFactory) Create() (*ent.User, error) {
	return f.client.User.Create().
		SetEmail(f.Email).
		SetUsername(f.Username).
		SetPassword(f.Password).
		SetCreatedAt(f.CreatedAt).
		SetUpdatedAt(f.UpdatedAt).
		AddPlanets(f.Edges.Planets...).
		Save(f.ctx)
}

// MustCreate returns the user struct, which is saved to the database.
// If an error occurs a panic is raised.
func (f *userFactory) MustCreate() *ent.User {
	u, err := f.Create()
	if err != nil {
		panic(err)
	}
	return u
}

// fakePassword generates a random hashed password. This password should
// not be used except for testing.
func fakePassword() string {
	password, err := bcrypt.GenerateFromPassword([]byte(faker.Internet().Password(6, 12)), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(password)
}
