package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/user"
	"github.com/stretchr/testify/assert"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestSignUpAccount(t *testing.T) {
	ts := withTestServer()
	defer ts.Close()

	generateVerifyToken = func() (string, error) {
		return "some_token", nil
	}

	data.GenerateEntity(0, 1, 2, 0, planet.PlanetTypeHabitable, "earth", "Earth")

	sent := false
	sendEmail = func(toAddress, toName, subject string, contents io.Reader) error {
		sent = true
		assert.Equal(t, "john@example.com", toAddress)
		assert.Equal(t, "john.doe", toName)
		assert.Equal(t, "Welcome to Empire and Foundation", subject)
		c, err := ioutil.ReadAll(contents)
		assert.NoError(t, err)
		assert.Contains(t, string(c), "john.doe")
		assert.Contains(t, string(c), html.EscapeString("/confirm_email?email=john%40example.com&token=some_token"))
		return nil
	}

	v := url.Values{}
	v.Add("csrf_token", "some_token")
	v.Add("email", "john@example.com")
	v.Add("username", "john.doe")
	v.Add("password", "secret01")
	v.Add("password_confirm", "secret01")

	res, err := http.Post(ts.URL+"/signup_account", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.True(t, sent)

	u, err := data.Client.User.Query().
		Where(user.UsernameEQ("john.doe")).
		First(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "john@example.com", u.Email)
	assert.Equal(t, "john.doe", u.Username)
	assert.Equal(t, "some_token", u.VerifyToken)
	p, err := u.QueryPlanets().All(context.Background())
	assert.NoError(t, err)
	assert.Len(t, p, 1)
	assert.Equal(t, "Earth", p[0].Name)
}

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
