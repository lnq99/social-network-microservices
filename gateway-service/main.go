package main

import (
	"app/config"
	"app/repository"
	"app/services"
	"app/util"
	"context"
	"fmt"
	"github.com/go-chi/cors"
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
	cfg config.Config
)

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)

	services.ProfilesServiceAddr = cfg.Service.ProfilesAddr
	services.PostsServiceAddr = cfg.Service.PostsAddr
	services.StatsServiceAddr = cfg.Service.StatsAddr
	services.TokenAuth = jwtauth.New("HS256", []byte(cfg.Auth.JwtSigningKey), nil)

	runDBMigration(cfg.Migration.Url, cfg.Db.Url)
}

func main() {
	db, err := util.NewPgxPool(cfg.Db.Url, context.Background())
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	defer db.Close()

	runDBMigration(cfg.Migration.Url, cfg.Db.Url)

	repo := repository.NewSqlRepository(db)

	r := chi.NewRouter()

	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(services.TokenAuth))

		// Handle valid / invalid tokens
		r.Use(jwtauth.Authenticator)

		//r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		//	_, claims, _ := jwtauth.FromContext(r.Context())
		//	w.Write([]byte(fmt.Sprintf("protected area. hi %+v", claims)))
		//})

		r.Route("/api/v1", func(r chi.Router) {
			//rProtected := r.Use(services.AuthMiddleware)

			s := services.ChiRouter{r}
			s.RegisterService(services.NewProfilesService())
			s.RegisterService(services.NewPostsService())
			s.RegisterService(services.NewNewsFeedService())
			s.RegisterService(services.NewSearchService())
			s.RegisterService(services.NewStatsService())
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/manage/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Route("/api/v1/auth", func(r chi.Router) {
			r.Use(jwtauth.Verifier(services.TokenAuth))
			s := services.ChiRouter{r}
			s.RegisterService(services.NewAuthService(repo))
		})

		//r.Route("/api/v1/feed", func(r chi.Router) {
		//	s := services.ChiRouter{r}
		//	s.RegisterService(services.NewNewsFeedService())
		//})

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
		log.Fatal(err, " - cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err, " - failed to run migrate up")
	}

	log.Println("db migrated successfully")
}
