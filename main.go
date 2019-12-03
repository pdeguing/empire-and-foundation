package main

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

func main() {
	info("Starting server...")

	// Public routes
	r := mux.NewRouter()
	files := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", files))

	r.HandleFunc("/", index)

	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", serveLogout)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signup_account", signupAccount).Methods("POST")
	r.HandleFunc("/authenticate", serveAuthenticate).Methods("POST")

	// Routes that require authentication
	rAuth := r.NewRoute().Subrouter()
	rAuth.HandleFunc("/dashboard", dashboard)
	rAuth.HandleFunc("/planet/up_metal_mine", upMetalMine)

	// Middleware
	csrfMiddleware := csrf.Protect(
		securecookie.GenerateRandomKey(32),
		csrf.FieldName("csrf_token"),
		csrf.CookieName("csrf_cookie"),
		csrf.Secure(false), // TODO: Remove this part once we support HTTPS.
		csrf.ErrorHandler(http.HandlerFunc(invalidCsrfToken)),
	)
	sessionMiddleware := sessionManager.LoadAndSave
	r.Use(
		csrfMiddleware,
		sessionMiddleware,
	)
	// Only apply the authentication middleware to the auth subrouter.
	rAuth.Use(
		authMiddleware,
	)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	info("Server started")
	err := server.ListenAndServe()
	if err != nil {
		danger(err, "An error occurred while running the server")
	}
}
