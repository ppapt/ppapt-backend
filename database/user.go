// The database package is our abstraction layer towards the databases.
//
// This file, user.go, contains functions to query information from the user
// table
package database

import (
	"github.com/ppapt/ppapt-backend/data"
	log "github.com/sirupsen/logrus"
)

// GetUser fetches user data from the database. It returns an user data
// structure or an error
// EMail is the EMail address (primary key) of the user to retrieve
func (db Database) GetUser(EMail string) (*data.User, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "GetUser"}).Logger
	logger.WithField("EMail", EMail).Debug("Fetch data")
	u := new(data.User)
	row,err := db.QueryRow("user_get", EMail)
	if err != nil {
		logger.WithField("action", "db.QueryRow").Error(err)
		return nil, err
	}
	err = row.Scan(&u.EMail, &u.Name, &u.Password, &u.Locked)
	if err != nil {
		logger.WithField("action", "row.Scan").Error(err)
		return nil, err
	}
	return u, nil
}

// UpdateUser updates the user with the given EMail address using the provided
// user data. It returns an error or nil in case of success.
// EMail is the current email address
// User is a data.User struct containing the updated data
func (db Database) UpdateUser(EMail string, User *data.User) error {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "UpdateUser"}).Logger
	logger.WithFields(log.Fields{
		"EMail":  User.EMail,
		"Name":   User.Name,
		"Locked": User.Locked,
	}).Debug("Execute Query")

	_, err := db.Exec("user_update", EMail, User.EMail, User.Name, User.Password, User.Locked)
	if err != nil {
		logger.WithField("action", "db.Exec").Error(err)
		return err
	}
	return nil
}

// AddUser adds a new user to the database. It returns nil on success or an error
// User is a data.User struct containing the EMail address, name, password and
// locked status of the new user
func (db Database) AddUser(User *data.User) error {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "AddUser"}).Logger
	logger.WithFields(log.Fields{
		"EMail":  User.EMail,
		"Name":   User.Name,
		"Locked": User.Locked,
	}).Debug("Execute Query")

	_, err := db.Exec("user_add", User.EMail, User.Name, User.Password, User.Locked)
	if err != nil {
		logger.WithField("action", "db.Exec").Error(err)
		return err
	}
	return nil
}
