// This package contains data structures used in several packages
package common

// UserListEntry represents one entry in a list of users
type UserListEntry struct {
	EMail  string `json:"email"`
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

// UserList is a list of users
type UserList []UserListEntry
