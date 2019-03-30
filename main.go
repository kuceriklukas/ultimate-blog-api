package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8008"

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {
	log.Println("Server started ...")

	author1 := *NewAuthor("John", "", "Incredible")
	author2 := *NewAuthor("Alice", "Raindeer", "Unicorn")

	posts = append(posts, *NewPost("First Post!", "The first first ultra post", author1))
	posts = append(posts, *NewPost("What's UltimateBlog about?", "It's about making blogs ulitmate!", author2))
	posts = append(posts, *NewPost("Understanding JavaScript closures", "It was always hard for me to understand closures, so I decided to actually learn what they are!", author2))

	router := mux.NewRouter()
	router.HandleFunc("/", getPosts).Methods("GET")

	log.Printf("Listening on %s \n", port)

	http.ListenAndServe(port, router)

}
