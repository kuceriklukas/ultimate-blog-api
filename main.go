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

	posts = append(posts, *NewPost("First Post!", "The first first ultra post"))
	posts = append(posts, *NewPost("What's UltimateBlog about?", "It's about making blogs ulitmate!"))
	posts = append(posts, *NewPost("Understanding JavaScript closures", "It was always hard for me to understand closures, so I decided to actually learn what they are!"))

	router := mux.NewRouter()
	router.HandleFunc("/", getPosts).Methods("GET")

	log.Printf("Listening on %s \n", port)

	http.ListenAndServe(port, router)

}
