package models

import (
	"errors"
	"fmt"
)

/*
  struct for user type
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

/*
GetUsers - Get all users
*/
func GetUsers() []*User {
	return users
}

/*
AddUser - Add the user by id
*/
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be set")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

/*
GetUserByID - Get the user by id
*/
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

/*
RemoveUserByID - Remove the user by id
*/
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id '%v' not found", id)
}

/*
UpdateUser - Update the user
*/
func UpdateUser(user User) (User, error) {
	for i, u := range users {
		if u.ID == user.ID {
			*users[i] = user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", user.ID)
}
