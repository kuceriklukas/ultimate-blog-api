package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kuceriklukas/ultimate-blog-api/models"
	"github.com/kuceriklukas/ultimate-blog-api/repositories"
)

func findPostById(posts []models.Post, id string) (*models.Post, error) {
	for _, post := range posts {
		if post.ID == id {
			return &post, nil
		}
	}

	return nil, errors.New("Couldn't find post")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.Posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	id, exists := r.URL.Query()["id"]
	if !exists || len(id[0]) < 1 {
		log.Println("Non-existing id parameter!")
		http.Error(w, "Non-existing id parameter!", http.StatusBadRequest)
		return
	}

	post, err := findPostById(repositories.Posts, id[0])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*post)
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
