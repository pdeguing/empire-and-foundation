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

	r.HandleFunc("/login", serveLogin)
	r.HandleFunc("/logout", serveLogout)
	r.HandleFunc("/signup", serveSignup)
	r.HandleFunc("/signup_account", serveSignupAccount).Methods("POST")
	r.HandleFunc("/authenticate", serveAuthenticate).Methods("POST")

	// Routes that require authentication
	rAuth := r.NewRoute().Subrouter()
	rAuth.HandleFunc("/dashboard", serveDashboard)
	rAuth.HandleFunc("/dashboard/cartography", serveCartography)
	rAuth.HandleFunc("/dashboard/fleetcontrol", serveFleetControl)
	rAuth.HandleFunc("/dashboard/technology", serveTechnology)
	rAuth.HandleFunc("/dashboard/diplomacy", serveDiplomacy)
	rAuth.HandleFunc("/dashboard/story", serveStory)
	rAuth.HandleFunc("/dashboard/wiki", serveWiki)
	rAuth.HandleFunc("/dashboard/news", serveNews)

	rPlanet := rAuth.PathPrefix("/planet/{planetNumber:[0-9]+}").Subrouter()
	rPlanet.HandleFunc("/", servePlanet)
	rPlanet.HandleFunc("/constructions", serveConstructions)
	rPlanet.HandleFunc("/factories", serveFactories)
	rPlanet.HandleFunc("/research", serveResearch)
	rPlanet.HandleFunc("/fleets", serveFleets)
	rPlanet.HandleFunc("/defenses", serveDefenses)
	rPlanet.HandleFunc("/metal-mine/upgrade", serveUpgradeMetalMine).Methods("POST")
	rPlanet.HandleFunc("/metal-mine/cancel", serveCancelMetalMine).Methods("POST")
	rPlanet.HandleFunc("/hydrogen-extractor/upgrade", serveUpgradeHydrogenExtractor).Methods("POST")
	rPlanet.HandleFunc("/hydrogen-extractor/cancel", serveCancelHydrogenExtractor).Methods("POST")
	rPlanet.HandleFunc("/silica-quarry/upgrade", serveUpgradeSilicaQuarry).Methods("POST")
	rPlanet.HandleFunc("/silica-quarry/cancel", serveCancelSilicaQuarry).Methods("POST")
	rPlanet.HandleFunc("/solar-plant/upgrade", serveUpgradeSolarPlant).Methods("POST")
	rPlanet.HandleFunc("/solar-plant/cancel", serveCancelSolarPlant).Methods("POST")
	rPlanet.HandleFunc("/housing-facilities/upgrade", serveUpgradeHousingFacilities).Methods("POST")
	rPlanet.HandleFunc("/housing-facilities/cancel", serveCancelHousingFacilities).Methods("POST")

	// Middleware
	csrfMiddleware := csrf.Protect(
		securecookie.GenerateRandomKey(32),
		csrf.FieldName("csrf_token"),
		csrf.CookieName("csrf_cookie"),
		csrf.Secure(false), // TODO: Remove this part once we support HTTPS.
		csrf.ErrorHandler(http.HandlerFunc(serveInvalidCsrfToken)),
	)
	sessionMiddleware := sessionManager.LoadAndSave
	r.Use(
		csrfMiddleware,
		sessionMiddleware,
		loadUserMiddleware,
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
