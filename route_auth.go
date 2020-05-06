package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// GET /signup
// Show the signup page
func serveSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "signup", nil, "layout", "public.navbar", "flash", "signup")
	forgetForm(r)
}

// POST /signup
// Create the user account
func serveSignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to parse form: %v", err)))
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 14)
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to encrypt password: %v", err)))
		return
	}

	err = data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		u, err := tx.User. // UserClient.
					Create().                             // User create builder.
					SetUsername(r.PostFormValue("name")). // Set field value.
					SetEmail(r.PostFormValue("email")).
					SetPassword(string(password)).
					Save(r.Context()) // Create and return.

		// TODO: Check availability of email address.
		// TODO: Validate inputs.
		if err != nil {
			return err
		}

		c, err := tx.Planet.Query().
			Where(
				planet.And(
					planet.PlanetTypeEQ(planet.PlanetTypeHabitable),
					planet.Not(planet.HasOwner()),
				),
			).
			Count(r.Context())

		if err != nil {
			return err
		}

		ra := rand.New(rand.NewSource(42))

		n := ra.Intn(c)

		p, err := tx.Planet.Query().
			Where(
				planet.And(
					planet.PlanetTypeEQ(planet.PlanetTypeHabitable),
					planet.Not(planet.HasOwner()),
				),
			).
			Offset(n).
			First(r.Context())

		_, err = p.Update().
			SetOwner(u).
			SetMetal(800).
			SetHydrogen(500).
			SetSilica(600).
			SetLastResourceUpdate(time.Now()).
			Save(r.Context())

		return err
	})
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to create a user account: %v", err)))
		return
	}

	flash(r, flashSuccess, "Your account has been created. You can log in now.")
	http.Redirect(w, r, "/", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func serveAuthenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u, err := data.Client.User.
		Query().
		Where(user.Email(r.PostFormValue("email"))).
		Only(r.Context())
	var nferr *ent.NotFoundError
	if errors.As(err, &nferr) {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/", 302)
		return
	}
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to check if user exists in database: %v", err)))
		return
	}
	ok, err := data.CheckPassword(u.Password, r.PostFormValue("password"))
	if err != nil {
		serveError(w, r, newInternalServerError(fmt.Errorf("unable to validate the users password: %v", err)))
		return
	}
	if !ok {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/", 302)
		return
	}

	authenticate(r, u)
	http.Redirect(w, r, "/dashboard", 302)
}

// GET /logout
// Logs the user out
func serveLogout(w http.ResponseWriter, r *http.Request) {
	logout(r)
	http.Redirect(w, r, "/", 302)
}
