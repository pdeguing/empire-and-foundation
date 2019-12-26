package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/urfave/cli/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		danger("Error loading .env file")
		os.Exit(1)
	}

	app := &cli.App{
		Name:  "empire-and-foundation",
		Usage: "server and maintenance tools",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "db-driver",
				Usage:    "use the 'mysql' or 'postgres' driver",
				EnvVars:  []string{"DB_DRIVER"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "db-user",
				Usage:    "connect to the database using `USERNAME`",
				Value:    "root",
				EnvVars:  []string{"DB_USERNAME"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "db-password",
				Usage:   "connect to the database using `PASSWORD`",
				EnvVars: []string{"DB_PASSWORD"},
			},
			&cli.StringFlag{
				Name:     "db-host",
				Usage:    "IP address or `HOSTNAME` on which the database is reachable",
				Value:    "localhost",
				EnvVars:  []string{"DB_HOST"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "db-port",
				Usage:    "connect to the database on `PORT`",
				Value:    5432,
				EnvVars:  []string{"DB_PORT"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "db-name",
				Usage:    "use `DATABASE`",
				Value:    "empire_and_foundation",
				EnvVars:  []string{"DB_DATABASE"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "db-debug",
				Usage:   "print SQL queries executed by the ent ORM",
				Value:   false,
				EnvVars: []string{"DB_DEBUG"},
			},
		},
		Before: func(c *cli.Context) error {
			d := c.String("db-driver")
			var connStr string
			switch d {
			case "mysql":
				connStr = mysqlConnString(c)
			case "postgres":
				connStr = postgresqlConnString(c)
			default:
				return fmt.Errorf("%q is not a supported driver", d)
			}
			err := data.InitDatabaseConnection(d, connStr, c.Bool("db-debug"))
			if err != nil {
				return err
			}

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s", "start"},
				Usage:   "start the webserver",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "run the server on `PORT`",
						Value:   8080,
						EnvVars: []string{"PORT"},
					},
				},
				Action: func(c *cli.Context) error {
					port := c.Int("port")
					info(fmt.Sprintf("Starting server on http://localhost:%d", port))
					initSessionManager(c.String("db-driver"))
					server := &http.Server{
						Addr:    ":" + strconv.Itoa(port),
						Handler: routes(),
					}
					return server.ListenAndServe()
				},
			},
			{
				Name:     "migrate",
				Category: "migrations",
				Usage:    "run the database migrations",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "drop-index",
						Usage: "drop indexes during migration",
						Value: false,
					},
					&cli.BoolFlag{
						Name:  "drop-column",
						Usage: "drop columns during migration",
						Value: false,
					},
				},
				Action: func(c *cli.Context) error {
					return data.Migrate(c.Context, data.Client, c.Bool("drop-index"), c.Bool("drop-column"))
				},
			},
			{
				Name:     "seed",
				Category: "migrations",
				Usage:    "seed the database with a generated world",
				Action: func(c *cli.Context) error {
					data.GenerateRegion(0)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		danger(err)
		os.Exit(1)
	}
}

func routes() http.Handler {
	// Public routes
	r := mux.NewRouter()
	files := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", files))

	r.HandleFunc("/", serveIndex)

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

	return r
}

// mysqlConnString uses the cli context to build a MySQL connection string.
func mysqlConnString(c *cli.Context) string {
	str := fmt.Sprintf("tcp(%s:%d)/%s?parseTime=true", c.String("db-host"), c.Int("db-port"), c.String("db-name"))
	u := c.String("db-user")
	p := c.String("db-password")
	if u != "" {
		if p == "" {
			str = fmt.Sprintf("%s@%s", u, str)
		} else {
			str = fmt.Sprintf("%s:%s@%s", u, p, str)
		}
	}
	return str
}

// postgresqlConnString uses the cli context to build a PostgreSQL connection string.
func postgresqlConnString(c *cli.Context) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.String("db-host"),
		c.Int("db-port"),
		c.String("db-user"),
		c.String("db-password"),
		c.String("db-name"),
	)
}
