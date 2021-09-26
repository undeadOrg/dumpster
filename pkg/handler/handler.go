package handler

import (
	"dumpster/pkg/storage"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// SavePayload - write out question to database
func (h *Handler) SavePayload(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	post := &storage.Payload{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ev.AddField("save.post", post.ID)

	// When do you convert the Request Object to a StorageObject?....
	// I'm going to cheat here... for now
	// TODO: Fix this for RequestObject conversion to Storage Object somehow?...
	queryStart := time.Now()

	err = h.Data.SaveObject(r.Context(), post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error"+err.Error())
	}
	ev.AddField("timers.db.response_insert_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// UpdateResponse - PUT question updated with additional answers
func (h *Handler) UpdateSocial(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	post := storage.Payload{}
	json.NewDecoder(r.Body).Decode(&post)

	queryStart := time.Now()
	_, err = h.Data.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not update question")
	}
	ev.AddField("timers.db.response_update_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Question Updated"})
}

// GetResponseID - Fetch a question from givin ID in url
func (h *Handler) GetSocialID(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	ev.AddField("response.id", id)

	queryStart := time.Now()
	result, err := h.Data.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not Fetch Response")
	}
	ev.AddField("timers.db.response_get_id_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, result)
}

// GetSocials - GET list of all questions/responses
func (h *Handler) GetSocials(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	queryStart := time.Now()
	results, err := h.Data.ListObjects(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}
	ev.AddField("timers.db.response_get_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, results)
}

// DeleteSocial - DELETE question id from questions
func (h *Handler) DeleteSocial(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	ev.AddField("response.id", id)

	queryStart := time.Now()
	err = h.Data.DeleteObject(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error Deleting ID")
	}
	ev.AddField("timers.db.response_delete_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
}

// respondwithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
