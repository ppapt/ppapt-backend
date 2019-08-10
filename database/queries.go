// The database package is our abstraction layer towards the databases.
//
// This file, queries.go, contains the SQL statements in the various database
// driver dialects
package database

// Query variant is a map of dialect->nativeSQL. It is used in QueryVatiants
type QueryVariant map[string]string

// QueryVariants is a map of short name->QueryVariant.
type QueryVariants map[string]QueryVariant

// Queries contains all queries used to talk to the database in the native
// dialect supported by the different database types.
var Queries = QueryVariants{
	"user_add": {
		"postgres": "INSERT INTO users (email, user_name, user_password, user_locked) VALUES ($1, $2, $3, $4);",
		"mysql":    "INSERT INTO users (email, user_name, user_password, user_locked) VALUES (?, ?, ?, ?);",
		"sqlite3":  "INSERT INTO users SET email=?, user_name=?, user_password=?, user_locked=?;",
	},
	"user_get": {
		"postgres": "SELECT * from users where email=$1;",
		"mysql":    "SELECT * from users where email=?;",
		"sqlite3":  "SELECT * from users where email=?;",
	},
	"user_update": {
		"postgres": "UPDATE users SET email=$1, user_name=$2, user_password=$3, user_locked=$4 where email=$5;",
		"mysql":    "UPDATE users SET email=?, user_name=?, user_password=?, user_locked=? where email=?;",
		"sqlite3":  "UPDATE users SET email=?, user_name=?, user_password=?, user_locked=? where email=?;",
	},
	"users_list": {
		"postgres": "SELECT email, user_name, user_locked FROM users ORDER BY email;",
		"mysql":    "SELECT email, user_name, user_locked FROM users ORDER BY email;",
		"sqlite3":  "SELECT email, user_name, user_locked FROM users ORDER BY email;",
	},
}
