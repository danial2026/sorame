package handler

import (
	"encoding/json"
	"net/http"
	"sorame/model"

	"github.com/gorilla/mux"
)

var linkRepo *model.LinkRepository

func SetLinkRepo(repo *model.LinkRepository) {
	linkRepo = repo
}

// InsertLink handles the creation of a new link
func InsertLink(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var link model.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the link
	if err := link.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the link in Redis
	linkUid, err := linkRepo.InsertLink(&link)
	if err != nil {
		http.Error(w, "Failed to insert link", http.StatusInternalServerError)
		return
	}

	// Create response with the link UID
	response := map[string]string{
		"share_id": linkUid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetLink handles retrieving a link by share ID
func GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shareID := vars["shareID"]

	// Get the link from Redis
	link, err := linkRepo.GetLink(shareID)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	// Redirect to the stored link with the original URL as param
	redirectURL := link.Data + "?shareUrl=" + "https://" + r.Host + "/link/" + shareID

	http.Redirect(w, r, redirectURL, http.StatusFound)
}
