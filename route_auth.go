package main

import (
	"database/sql"
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
)

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "login.layout", "public.navbar", "flash", "login")
	forgetForm(r)
}

// GET /signup
// Show the signup page
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "login.layout", "public.navbar", "flash", "signup")
	forgetForm(r)
}

// POST /signup
// Create the user account
func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		internalServerError(w, r, err, "Cannot parse form")
		return
	}
	user := data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	// TODO: Check availability of email address.
	// TODO: Validate inputs.
	if err := user.Create(); err != nil {
		internalServerError(w, r, err, "Cannot create user")
		return
	}
	flash(r, flashSuccess, "Your account has been created. You can log in now.")
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func serveAuthenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err == sql.ErrNoRows {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/login", 302)
		return
	}
	if err != nil {
		internalServerError(w, r, err, "Cannot retrieve user by email")
		return
	}
	ok, err := user.CheckPassword(r.PostFormValue("password"))
	if err != nil {
		internalServerError(w, r, err, "Cannot check user's password")
		return
	}
	if !ok {
		flash(r, flashDanger, "The username or password you have entered is invalid.")
		rememberForm(r)
		http.Redirect(w, r, "/login", 302)
		return
	}

	authenticate(r, &user)
	http.Redirect(w, r, "/dashboard", 302)
}

// GET /logout
// Logs the user out
func serveLogout(w http.ResponseWriter, r *http.Request) {
	logout(r)
	http.Redirect(w, r, "/", 302)
}
