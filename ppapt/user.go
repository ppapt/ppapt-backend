// The data package defines ppapt data structures and functions to handle them
//
// This file user.go contains the data structures to handle user related data
// (database structure as well as server responses)

// User is the server side data structure for the database user
package ppapt

import (
	"errors"
	"github.com/ppapt/ppapt-backend/common"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"regexp"
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

// HashPassword uses bcrypt to hash the given password. On success, the hash is
// returned, otherwise an error is returned.
// Cleartext is the cleartext password
func (p *Ppapt) HashPassword(Cleartext string) (string, error) {
	c := []byte(Cleartext)
	h, err := bcrypt.GenerateFromPassword(c, 10)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

// GetUser retrieves an user object from the database
// EMail
func (p *Ppapt) GetUser(EMail string) (*User, error) {
	logger := log.WithFields(log.Fields{
		"context": "ppapt",
		"func":    "FetUser"})
	logger.WithField("EMail", EMail).Debug("Get user data")
	e, n, pass, l, err := p.db.GetUser(EMail)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	u := &User{
		EMail:    e,
		Name:     n,
		Password: pass,
		Locked:   l,
	}
	return u, nil
}

// Login validates the password for the given email and generates or updates a
// token. It returns a pointer to the user object and the token or an error.
// EMail is the users email address
// Password the password to be validated.
func (p *Ppapt) Login(EMail string, Password string) (*User, string, error) {
	logger := log.WithFields(log.Fields{
		"context": "ppapt",
		"func":    "Login",
		"EMail":   EMail,
	})
	logger.Debug("Get user data")
	u, err := p.GetUser(EMail)
	if err != nil {
		logger.Warn("Could not retrieve user data")
		return nil, "", err
	}
	if !u.PasswordIsValid(Password) {
		logger.Warn("Invalid password")
		return nil, "", err
	}
	t := p.session.AddSession(EMail)
	return u, t, nil
}

func (p *Ppapt) NewUser(EMail string, Name string, CleartextPassword string, Locked bool) (*User, error) {
	logger := log.WithFields(log.Fields{
		"context": "ppapt",
		"func":    "NewUser",
		"EMail":   EMail,
		"Name":    Name,
		"Locked":  Locked,
	})
	logger.Debug("New user")
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(EMail) {
		e := "EMail does not match validation regex"
		logger.Error(e)
		return nil, errors.New(e)
	}
	if Name == "" {
		e := "Name must not be empty"
		logger.Error(e)
		return nil, errors.New(e)
	}
	pass, err := p.HashPassword(CleartextPassword)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	err = p.db.AddUser(EMail, Name, pass, Locked)
	if err != nil {
		return nil, err
	}
	u := &User{
		EMail:    EMail,
		Name:     Name,
		Password: pass,
		Locked:   Locked,
	}
	return u, nil
}

func (p *Ppapt) ListUsers() (common.UserList, error) {
	logger := log.WithFields(log.Fields{
		"context": "ppapt",
		"func":    "ListUsers",
	})
	logger.Debug("List users")
	u, err := p.db.ListUsers()
	return u, err
}
