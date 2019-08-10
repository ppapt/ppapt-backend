// The database package is our abstraction layer towards the databases.
//
// This file, user.go, contains functions to query information from the user
// table
package database

import (
	"github.com/ppapt/ppapt-backend/common"
	log "github.com/sirupsen/logrus"
)

// GetUser fetches user data from the database. It returns an user data
// structure or an error
// EMail is the EMail address (primary key) of the user to retrieve
func (db Database) GetUser(EMail string) (string, string, string, bool, error) {
	var REMail string
	var Name string
	var Password string
	var Locked bool
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "GetUser"})
	logger.WithField("EMail", EMail).Debug("Fetch data")
	row, err := db.QueryRow("user_get", EMail)
	if err != nil {
		logger.WithField("action", "db.QueryRow").Error(err)
		return "", "", "", false, err
	}
	err = row.Scan(REMail, Name, Password, Locked)
	if err != nil {
		logger.WithField("action", "row.Scan").Error(err)
		return "", "", "", false, err
	}
	return REMail, Name, Password, Locked, nil
}

// UpdateUser updates the user with the given EMail address using the provided
// user data. It returns an error or nil in case of success.
// EMail is the current email address
// NewEMail is the new EMail address for the user, leaving empty means no change
// Name is the users name
// Password is the hashed password
// Locked is true if the user is not allowed to log in
func (db Database) UpdateUser(EMail string, NewEmail string, Name string, Password string, Locked bool) error {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "UpdateUser"})
	logger.WithFields(log.Fields{
		"EMail":    EMail,
		"NewEmail": NewEmail,
		"Name":     Name,
		"Locked":   Locked,
	}).Debug("Execute Query")
	if NewEmail == "" {
		NewEmail = EMail
	}
	_, err := db.Exec("user_update", EMail, Name, Password, Locked, NewEmail)
	if err != nil {
		logger.WithField("action", "db.Exec").Error(err)
		return err
	}
	return nil
}

// AddUser adds a new user to the database. It returns nil on success or an error
// User is a data.User struct containing the EMail address, name, password and
// locked status of the new user
func (db Database) AddUser(EMail string, Name string, Password string, Locked bool) error {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "AddUser"})
	logger.WithFields(log.Fields{
		"EMail":  EMail,
		"Name":   Name,
		"Locked": Locked,
	}).Debug("Execute Query")

	_, err := db.Exec("user_add", EMail, Name, Password, Locked)
	if err != nil {
		logger.WithField("action", "db.Exec").Error(err)
		return err
	}
	return nil
}

func (db Database) ListUsers() (common.UserList, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "ListUsers"})
	rows, err := db.Query("users_list")
	if err != nil {
		logger.WithField("action", "db.Query").Error(err)
		return nil, err
	}
	u := make(common.UserList, 0)
	for rows.Next() {
		var e string
		var n string
		var l bool
		if err := rows.Scan(&e, &n, &l); err != nil {
			logger.WithField("action", "rows.Scan").Error(err)
			return nil, err
		}
		entry := common.UserListEntry{
			EMail:  e,
			Name:   n,
			Locked: l,
		}
		u = append(u, entry)
	}
	if err := rows.Err(); err != nil {
		logger.WithField("action", "rows.Next").Error(err)
		return nil, err
	}
	return u, nil
}
