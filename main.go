package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	flag.Parse()

	//c := config.New()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Setup Connection timeouts
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Connect to database
	//fmt.Printf("Connecting to Database\n")
	//qrepo, err := db.NewQuestionsRepo(ctx, c.URI, c.DB)
	//if err != nil {
		// Implement better health checking/retry here or in lib
	//	log.Fatalf("Cannot set up Database: %v", err)
	//}

	//qset, err := db.NewQuestionSetRepo(ctx, c.URI, "question_set")
	//if err != nil {
	//	log.Fatalf("Cannot set up Database: %v", err)
	//}

	//qhandler := questions.NewHandler(qrepo, qset)

	// Setup Router
	router := Router()

	router.Get("/", index)

	//router.Route("/api/v1", func(r chi.Router) {
	//	r.Mount("/questions", questions.Router(qhandler))
	//})
	fmt.Printf("Starting up Webserver\n")
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
