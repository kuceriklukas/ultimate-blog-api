package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuceriklukas/ultimate-blog-api/models"
	"github.com/kuceriklukas/ultimate-blog-api/repositories"
	"github.com/kuceriklukas/ultimate-blog-api/handlers"
)

const port = ":8008"

func prepareTempData() {
	author1 := *models.NewAuthor("John", "", "Incredibulous")
	author2 := *models.NewAuthor("Alice", "Starborn", "Unicorn")
	author3 := *models.NewAuthor("Luk", "A", "Chu")
	repositories.Authors = append(repositories.Authors, author1) 
	repositories.Authors = append(repositories.Authors, author2) 
	repositories.Authors = append(repositories.Authors, author3) 

	repositories.CurrentAuthor = author3

	repositories.Posts = append(repositories.Posts, *models.NewPost("First Post!", "The first first ultra post. Like, actually the first", author1))
	repositories.Posts = append(repositories.Posts, *models.NewPost("What's UltimateBlog about?", "It's about making blogs ulitmate! Waaaaa", author2))
	repositories.Posts = append(repositories.Posts, *models.NewPost("Understanding JavaScript closures", "It was always hard for me to understand closures, so I decided to actually learn what they are! Surprise!", author2))
}

func main() {
	log.Println("Server started ...")

	prepareTempData()
	log.Println("Data have been prepared ...")

	router := mux.NewRouter()
	router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	router.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	router.HandleFunc("/create-post", handlers.CreatePost).Methods("POST")

	log.Printf("Listening on %s \n", port)

	http.ListenAndServe(port, router)
}
