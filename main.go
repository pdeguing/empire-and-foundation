package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func main() {
	info("Starting server...")
	r := mux.NewRouter()
	files := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", files))

	r.HandleFunc("/", index)

	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signup_account", signupAccount).Methods("POST")
	r.HandleFunc("/authenticate", authenticate).Methods("POST")

	r.HandleFunc("/dashboard", dashboard)

	r.HandleFunc("/planet/up_metal_mine", upMetalMine)

	csrfWrapper := csrf.Protect(
		securecookie.GenerateRandomKey(32),
		csrf.FieldName("csrf_token"),
		csrf.CookieName("csrf_cookie"),
		csrf.Secure(false), // TODO: Remove this part once we support HTTPS.
		csrf.ErrorHandler(http.HandlerFunc(invalidCsrfToken)),
	)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: csrfWrapper(r),
	}
	info("Server started")
	server.ListenAndServe()
}
