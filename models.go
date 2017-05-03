package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Passport holds passport data
type Passport struct {
	ID           string    `json:"id"`
	DateOfIssue  time.Time `json:"dateOfIssue"`
	DateOfExpiry time.Time `json:"dateOfExpiry"`
	Authority    string    `json:"authority"`
	UserID       int       `json:"userId"`
}

// User holds personal user information
type User struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	FirstName       string        `json:"firstName"`
	LastName        string        `json:"lastName"`
	DateOfBirth     time.Time     `json:"dateOfBirth"`
	LocationOfBirth string        `json:"locationOfBirth"`
}
