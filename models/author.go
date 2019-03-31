package models

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

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
