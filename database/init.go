// The database package is our abstraction layer towards the databases.
//
// This file, init.go, contains the database connection routines and functions
// used in other functions
package database

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
)

// The Database structure is used to store the database, its context, the type
// and DSN for future reference. It is returned by the NewDatabase function.
type Database struct {
	db     *sql.DB
	ctx    context.Context
	DBType string
	DSN    string
}

// Create a new Database object based on the given database type and DSN
// DBType is one of "postgres", "mysql" or "sqlite3" and defines the type of database to connect to
// DSN must be a valid DSN for connecting to the database
func NewDatabase(DBType string, DSN string) (*Database, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "NewDatabase"}).Logger
	conn := new(Database)
	db, err := sql.Open(DBType, DSN)
	if err != nil {
		logger.WithField("action", "open").Fatal(err)
		return nil, err
	}
	conn.db = db
	conn.DBType = DBType
	conn.DSN = DSN

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	conn.ctx = ctx
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	if err = db.PingContext(ctx); err != nil {
		log.WithField("action", "ping").Fatal(err)
		return nil, err
	}
	return conn, nil
}

// BuildDSN creates a valid DSN for connecting to a database, based on the
// values of its parameters. It returns a string containing a valid DSN or
// an error, if the database type is not supported.
// DBType is one of "postgres", "mysql" or "sqlite3", if another type is provided, the function returns an error
// User is the username to connect to the database
// Password is the password for the database user
// DBName is the name of the database on the DBMS
// Server is the name of the database server
// Port is the port number, the database server is listening on
func BuildDSN(DBType string, User string, Password string, DBName string, Server string, Port int) (string, error) {
	switch DBType {
	case "mysql":
		return User + ":" + Password + "@mysql(" + Server + ":" + strconv.Itoa(Port) + ")/" + DBName, nil
	case "postgres":
		return "postgres://" + User + ":" + Password + "@" + Server + ":" + strconv.Itoa(Port) + "/" + DBType, nil
	case "sqlite3":
		//#ToDo: Add sqlite3 dsn handling
		return "", nil
	default:
		return "", errors.New("Unsupported database type " + DBType)
	}
}

// GetNativeSQL fetches the SQL statement for the given name in the SQL dialect
// of the databases DBType. An error is returned if the StatementName is unknown.
// StatementName is the name for that query (see queries.go)
func (db Database) GetNativeSQL(StatementName string) (string, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "GetNativeSQL"}).Logger
	if q, ok := Queries[StatementName][db.DBType]; ok {
		return q, nil
	} else {
		e := "Invalid statement name " + StatementName
		logger.WithFields(log.Fields{
			"context":       "database",
			"func":          "GetNativeSQL",
			"statementname": StatementName,
		}).Error(e)
		return "", errors.New(e)
	}
}

// Query runs a query (usually a select) with multiple rows as expected result.
// It returns a sql.Rows pointer or an error message returned by the database.
func (db Database) Query(StatementName string, Args ...interface{}) (*sql.Rows, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "Query"}).Logger
	SQL, err := db.GetNativeSQL(StatementName)
	if err != nil {
		return nil, err
	}
	logger.WithField("query", SQL).Debug("Execute query")
	rows, err := db.db.QueryContext(db.ctx, SQL, Args)
	if err != nil {
		log.WithField("action", "querycontext").Error(err)
		return nil, err
	}
	return rows, nil
}

// QueryRow runs a query (usually a select) with a single row as expected result.
// It returns a sql.Row pointer. Error handling is deferred until the scan of
// the row.
func (db Database) QueryRow(StatementName string, Args ...interface{}) (*sql.Row, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "QueryRow"}).Logger
	SQL, err := db.GetNativeSQL(StatementName)
	if err != nil {
		return nil, err
	}
	logger.WithField("query", SQL).Debug("Execute query")
	row := db.db.QueryRowContext(db.ctx, SQL, Args)
	return row, nil
}

// Exec runs a query (usually an instert, update or delete) which does not expect
// to get data back. It returns a sql.Result or an error message.
func (db Database) Exec(StatementName string, Args ...interface{}) (sql.Result, error) {
	logger := log.WithFields(log.Fields{
		"context": "database",
		"func":    "Exec"}).Logger
	SQL, err := db.GetNativeSQL(StatementName)
	if err != nil {
		return nil, err
	}
	logger.WithField("query", SQL).Debug("Execute query")
	result, err := db.db.ExecContext(db.ctx, SQL, Args)
	if err != nil {
		log.WithField("action", "execcontext").Error(err)
		return result, err
	}
	return result, nil
}
