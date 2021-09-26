package main

import (
	"context"
	"dumpster/pkg/config"
	"dumpster/pkg/db"
	"dumpster/pkg/handler"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Debug - Enable debug logging
var Debug = flag.Bool("debug", false, "Enable Debug Logging")

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5, ""),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	// Initialize Basic Ping functionality
	router.Get("/ping", ping)

	return router
}

func main() {
	var port = "5000"
	flag.StringVar(&port, "port", port, "Port")
	flag.Parse()

	// Initialize Basic Config
	conf := config.New()

	// Setup Connection timeouts
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Connect to database
	log.Printf("Connecting to Database\n")
	s, err := db.NewDumpsterRepo(ctx, conf)
	if err != nil {
		// Implement better health checking/retry here or in lib
		log.Fatalf("Cannot set up Database: %v", err)
	}

	handlers := handler.NewHandler(s)

	// Setup Router
	router := Router()

	router.Get("/", index)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/social", handler.Router(handlers))
	})
	log.Printf("Starting up Webserver\n")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello world"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("."))
}
