package handler

import (
	"dumpster/pkg/storage"
	"net/http"

	"github.com/go-chi/chi"
)

// Handler - Context for Handling of Questions routes
type Handler struct {
	Data storage.DumpsterData
}

// NewHandler - Initialize Handler
func NewHandler(d storage.DumpsterData) *Handler {
	return &Handler{
		Data: d,
	}
}

// Router - A completely separate router the questions data storage handle
func Router(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(authHandler)

	r.Get("/", HoneycombMiddleware(h.GetSocials))
	r.Get("/{id:[0-9a-f]+}", HoneycombMiddleware(h.GetSocialID))
	r.Post("/", HoneycombMiddleware(h.SavePayload))
	r.Put("/{id:[0-9a-f]+}", HoneycombMiddleware(h.UpdateSocial))
	r.Delete("/{id:[0-9a-f]+}", HoneycombMiddleware(h.DeleteSocial))

	return r
}
