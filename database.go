package main

import (
	"log"

	"github.com/palantir/stacktrace"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//"gopkg.in/mgo.v2/bson"

// DataStorer defines all the database operations
type IDao interface {
	ListUsers() ([]User, error)
	GetUser(uid string) (User, error)
	AddUser(u User) (User, error)
	DeleteUser(uid string) error
	//UpdateUser(u User) (User, error)
}

// MockDB will hold the connection and key db info
type Dao struct {
	DB      *mgo.Database
	Session *mgo.Session
	User    *mgo.Collection
}

// AddUser adds a User JSON document, returns the JSON document with the generated id
func (dao *Dao) AddUser(u User) (User, error) {

	err := dao.User.Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u, nil
}

// ListUsers returns a list of JSON documents
func (dao *Dao) ListUsers() ([]User, error) {
	var result []User

	err := dao.User.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

// GetUser returns a single JSON document
func (dao *Dao) GetUser(uid string) (User, error) {

	result := User{}
	err := dao.User.FindId(bson.ObjectIdHex(uid)).One(&result)
	if err != nil {
		log.Fatal(err)
		return User{}, stacktrace.NewError("Failure trying to retrieve user")
	}
	return result, nil
}

// DeleteUser deletes a user
func (dao *Dao) DeleteUser(uid string) error {

	err := dao.User.RemoveId(bson.ObjectIdHex(uid))
	if err != nil {
		return stacktrace.NewError("Failure trying to delete user")
	}
	return nil
}

/*
// UpdateUser updates an existing user
func (dao *Dao) UpdateUser(u User) (User, error) {
	id := u.ID
	user, err := dao.GetUser(id)

	if err != nil {
		log.Fatal(err)
		return u, stacktrace.NewError("Failure trying to update user")
	}


	dao.User.
	_, ok := db.UserList[id]
	if !ok {
		return u, stacktrace.NewError("Failure trying to update user")
	}
	db.UserList[id] = u
	return db.UserList[id], nil
}
*/
