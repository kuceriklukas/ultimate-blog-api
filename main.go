package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuceriklukas/ultimate-blog-api/models"
)

const port = ":8008"

var authors []models.Author
var posts []models.Post
var currentAuthor models.Author

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func createPost(w http.ResponseWriter, r *http.Request) {
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

	newPost := *models.NewPost(title[0], text[0], currentAuthor)
	posts = append(posts, newPost)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func main() {
	log.Println("Server started ...")

	author1 := *models.NewAuthor("John", "", "Incredibulous")
	author2 := *models.NewAuthor("Alice", "Starborn", "Unicorn")
	author3 := *models.NewAuthor("Luk", "A", "Chu")
	authors = append(authors, author1) 
	authors = append(authors, author2) 
	authors = append(authors, author3) 

	currentAuthor = author3

	posts = append(posts, *models.NewPost("First Post!", "The first first ultra post", author1))
	posts = append(posts, *models.NewPost("What's UltimateBlog about?", "It's about making blogs ulitmate!", author2))
	posts = append(posts, *models.NewPost("Understanding JavaScript closures", "It was always hard for me to understand closures, so I decided to actually learn what they are!", author2))

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/authors", getAuthors).Methods("GET")
	router.HandleFunc("/create-post", createPost).Methods("POST")

	log.Printf("Listening on %s \n", port)

	http.ListenAndServe(port, router)

}
