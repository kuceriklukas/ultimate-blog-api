package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kuceriklukas/ultimate-blog-api/repositories"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.Authors)
}