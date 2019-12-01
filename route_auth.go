package main

import (
	"database/sql"
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
)

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplatesFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

// GET /signup
// Show the signup page
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
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
	if err := user.Create(); err != nil {
		internalServerError(w, r, err, "Cannot create user")
		return
	}
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err == sql.ErrNoRows {
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
		http.Redirect(w, r, "/login", 302)
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		internalServerError(w, r, err, "Cannot create session")
		return
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	info("Successfully logged in, redirecting to root path...")
	http.Redirect(w, r, "/", 302)
}

// GET /logout
// Logs the user out
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
