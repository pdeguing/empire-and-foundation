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

	pf := data.NewPlanetFactory()
	pf.PlanetType = planet.PlanetTypeHabitable
	pf.Name = "Earth"
	pf.MustCreate()

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

func TestConfirmEmail(t *testing.T) {
	ts := withTestServer()
	defer ts.Close()

	uf := data.NewUserFactory()
	uf.Email = "john@example.com"
	uf.VerifyToken = "some_token"
	u := uf.MustCreate()

	res, err := http.Get(ts.URL+"/confirm_email?email=john%40example.com&token=some_token")
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	u, err = data.Client.User.Get(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.Empty(t, u.VerifyToken)
}

func TestAuthenticate_cannotAuthenticateIfAccountIsNotVerified(t *testing.T) {
	ts := withTestServer()
	defer ts.Close()

	uf := data.NewUserFactory()
	uf.Email = "john@example.com"
	uf.Password = "secret01"
	uf.VerifyToken = "some_token"
	uf.MustCreate()

	v := url.Values{}
	v.Add("csrf_token", "some_token")
	v.Add("email", "john@example.com")
	v.Add("password", "secret01")

	res, err := newTestClient().Post(ts.URL+"/authenticate", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, res.Request.URL.Path,"/")
	c, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(c), "The username or password you have entered is invalid.")
}