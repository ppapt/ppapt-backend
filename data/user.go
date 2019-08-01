// The data package defines data structures and functions to handle them
//
// This file user.go contains the data structures to handle user related data
// (database structure as well as server responses)

// User is the server side data structure for the database user
//`liquibase`
package data

import (
	"golang.org/x/crypto/bcrypt"
)

// User is the golang user object matching the user table in the database
type User struct {
	EMail    string `json: "email" liquibase:"email         varchar(128) not null" liquibase_table:"user" liquibase_key:"primary"`
	Name     string `json: "name" liquibase:"user_name     varchar(64) not null" liquibase_table:"user"`
	Password string `liquibase:"user_password varchar(60) not null" liquibase_table:"user"`
	Locked   bool   `json:"locked" liquibase:"user_locked   boolean not null" liquibase_table:"user"`
}

// LoginData is the data sent by a web caller to login or change his password
type LoginData struct {
	EMail       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// PasswordIsValid uses bcrypt to validate the given Password against the users
// hashed password.
// Password is the cleartext password. It is compared against the hashed
// password stored in the user object/database.
func (u User) PasswordIsValid(Password string) bool {
	p := []byte(Password)
	h := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(h, p)
	return (err == nil)
}
