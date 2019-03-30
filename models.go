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

// Author type
type Author struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
}

// NewAuthor - Author's constructor
func NewAuthor(firstName string, middleName string, lastName string) *Author {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Couldn't make a uuid when creating a new Author")
	}

	return &Author{ID: uid.String(), FirstName: firstName, MiddleName: middleName, LastName: lastName}
}
