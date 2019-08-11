// This package contains data structures used in several packages
package common

// LoginData is the data sent by a web caller to login or change his password
type LoginData struct {
	EMail       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// UserListEntry represents one entry in a list of users
type UserListEntry struct {
	EMail  string `json:"email"`
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

// UserList is a list of users
type UserList []UserListEntry

type Message struct {
	Message string `json:"message"`
}