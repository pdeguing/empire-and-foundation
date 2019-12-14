package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// GET /login
// Show the login page
func serveLogin(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "login", nil, "login.layout", "public.navbar", "flash", "login")
	forgetForm(r)
}

// GET /signup
// Show the signup page
func serveSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "signup", nil, "login.layout", "public.navbar", "flash", "signup")
	forgetForm(r)
}

// POST /signup
// Create the user account
func serveSignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot parse form")
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 14)
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot encrypt password")
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

		_, err = tx.Planet.
			Create().
			SetOwner(u).
			Save(r.Context())
		return err
	})
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot create user or planet")
		return
	}

	flash(r, flashSuccess, "Your account has been created. You can log in now.")
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func serveAuthenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u, err := data.Client.User.
		Query().
		Where(user.Email(r.PostFormValue("email"))).
		Only(r.Context())
	if err != nil {
		flash(r, flashDanger, "The username you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/login", 302)
		return
	}
	ok, err := data.CheckPassword(u.Password, r.PostFormValue("password"))
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot check user's password")
		return
	}
	if !ok {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/login", 302)
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
