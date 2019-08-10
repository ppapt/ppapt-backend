// The data package defines ppapt data structures and functions to handle them
//
// This file main.go contains the global variables and structures
// (database structure as well as server responses)

package ppapt

import (
	"github.com/ppapt/ppapt-backend/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// This is the main ppapt object bringing together al components
type Ppapt struct {
	db      *database.Database
	session Session
}

func NewPpapt() (*Ppapt, error) {
	ppapt := new(Ppapt)

	DBType := viper.GetString("database-type")
	DSN, err := database.BuildDSN(
		DBType,
		viper.GetString("database-user"),
		viper.GetString("database-password"),
		viper.GetString("database-name"),
		viper.GetString("database-server"),
		viper.GetInt("database-port"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	DB, err := database.NewDatabase(DBType, DSN)
	if err != nil {
		return nil, err
	}
	ppapt.db = DB
	return ppapt, nil
}
