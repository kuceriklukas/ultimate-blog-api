package main

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

// Post type
type Post struct {
	ID          string
	Title       string
	Description string
}

// NewPost - Post's constructor
func NewPost(title string, description string) *Post {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Can't make a uid!")
	}

	return &Post{ID: uid.String(), Title: title, Description: description}
}
