package main

import (
	"app/config"
	"app/services"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	cfg       *config.Config
	tokenAuth *jwtauth.JWTAuth
)

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)

	tokenAuth = jwtauth.New("HS256", []byte(cfg.Auth.JwtSigningKey), nil) // replace with secret key

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	r := chi.NewRouter()

	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		//r.Use(jwtauth.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})

		r.Route("/api/v1", func(r chi.Router) {
			//rProtected := r.Use(services.AuthMiddleware)

			s := services.ChiRouter{r}
			s.RegisterService(services.NewProfilesService(cfg.Service.ProfilesAddr))
			s.RegisterService(services.NewPostsService(cfg.Service.PostsAddr))
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/manage/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Route("/api/v1/auth", func(r chi.Router) {
			s := services.ChiRouter{r}
			s.RegisterService(services.NewAuthService())
		})

		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("route does not exist"))
		})

		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("method is not valid"))
		})
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s:\t'%s' has %d middlewares\n", method, route, middlewares)
		return nil
	})

	http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, r)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal(err, "cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err, "failed to run migrate up")
	}

	log.Println("db migrated successfully")
}
