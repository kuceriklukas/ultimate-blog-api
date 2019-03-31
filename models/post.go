package models

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

// Post type
type Post struct {
	ID          string
	Title       string
	Description string
	Author      Author
}

// NewPost - Post's constructor
func NewPost(title string, description string, author Author) *Post {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Couldn't make a uuid when creating a new Post")
	}

	return &Post{ID: uid.String(), Title: title, Description: description, Author: author}
}
