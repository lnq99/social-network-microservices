package util

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Run()
	SetupRouter()
}

type BaseServer struct {
	server *http.Server
}

type Config struct {
	Addr         string
	Handler      http.Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Addr:         ":8080",
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 8 * time.Second,
	}
}

func (s *BaseServer) InitHttpServer(cfg *Config) {
	s.server = &http.Server{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Handler:      cfg.Handler,
	}
}

func (s *BaseServer) Run() {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
