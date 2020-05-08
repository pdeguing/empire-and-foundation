package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendSignupEmail(t *testing.T) {
	t.SkipNow()

	data.WithTestDatabase()
	godotenv.Load()

	verifyToken, err := generateRandomString(20)
	assert.NoError(t, err)
	fmt.Println(verifyToken)

	u, err := data.Client.User.
		Create().
		SetUsername("John Doe").
		SetEmail("your@email.here"). // Change this to your email before running the test
		SetPassword("").
		SetVerifyToken(verifyToken).
		Save(context.Background())
	assert.NoError(t, err)

	assert.NoError(t, sendSignupEmail(u))
}
