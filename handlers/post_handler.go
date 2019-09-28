package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kuceriklukas/ultimate-blog-api/models"
	"github.com/kuceriklukas/ultimate-blog-api/repositories"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.Posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	title, exists := r.URL.Query()["title"]
	if !exists || len(title[0]) < 1 {
		log.Println("Non-existing title parameter!")
		http.Error(w, "Non-existing title parameter!", http.StatusBadRequest)
		return
	}

	text, exists := r.URL.Query()["text"]
	if !exists || len(text[0]) < 1 {
		log.Println("Non-existing text parameter!")
		http.Error(w, "Non-existing text parameter!", http.StatusBadRequest) 
		return
	}

	newPost := *models.NewPost(title[0], text[0], repositories.CurrentAuthor)
	repositories.Posts = append(repositories.Posts, newPost)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}