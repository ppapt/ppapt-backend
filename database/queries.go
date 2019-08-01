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
		"postgres": "INSERT INTO users SET email=$1, user_name=$2, user_password=$3, user_locked=$4",
		"mysql":    "INSERT INTO users SET email=?, user_name=?, user_password=?, user_locked=?",
		"sqlite3":  "INSERT INTO users SET email=?, user_name=?, user_password=?, user_locked=?",
	},
	"user_get": {
		"postgres": "SELECT * from users where email=$1",
		"mysql":    "SELECT * from users where email=?",
		"sqlite3":  "SELECT * from users where email=?",
	},
	"user_update": {
		"postgres": "UPDATE users SET email=$2, user_name=$3, user_password=$4, user_locked=$5 where email=$1",
		"mysql":    "UPDATE users SET email=?, user_name=?, user_password=?, user_locked=? where email=?",
		"sqlite3":  "UPDATE users SET email=?, user_name=?, user_password=?, user_locked=? where email=?",
	},
}
