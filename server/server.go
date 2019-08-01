// server is the REST API and webserver component of the ppapt-backend
//
// This file, server.go, contains the main router and functions used throughout
// the package.
package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ppapt/ppapt-backend/data"
	"github.com/ppapt/ppapt-backend/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

// DB stores the database connection established when starting a router
var DB *database.Database

// Session keeps all logged in users.
//#ToDo: This is currently not thread-safe, make this a separate object
var Session data.Session

// Router opens a web server and routes calls to the static web server or the
// various API functions
func Router() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/login", ApiLogin)
	//	router.PUT("/api/*path", ApiPUT)
	//	router.POST("/api/*path", ApiPOST)
	//	router.DELETE("/api/*path", ApiDELETE)
	log.Debug("Router initialized")
	router.ServeFiles("/static/*filepath", http.Dir("static/"))

	CertificateFile := viper.GetString("certificate-file")
	KeyFile := viper.GetString("key-file")
	Port := viper.GetInt("port")

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
		os.Exit(10)
	}
	DB, err = database.NewDatabase(DBType, DSN)
	if err != nil {
		log.Fatal(err)
		os.Exit(10)
	}

	if _, err := os.Stat(CertificateFile); err == nil {
		if _, err := os.Stat(KeyFile); os.IsNotExist(err) {
			log.Error("Found " + CertificateFile + " but missing " + KeyFile)
			os.Exit(10)
		}
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(Port), CertificateFile, KeyFile, router))

	} else {
		log.Fatal(CertificateFile + " not found")
		os.Exit(10)
	}
}

// RequestLogger logs a http request
// R is the http request
// Message the log message to be output at Info log level
func RequestLogger(R *http.Request, Message string) {
	log.WithFields(log.Fields{
		"method": R.Method,
		"url":    R.URL.String(),
		"remote": R.RemoteAddr,
		"proto":  R.Proto,
	}).Info(Message)
}

// sendJSON sends a JSON response out of the provided data
func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	j, err := json.Marshal(data)
	if err != nil {
		log.WithFields(log.Fields{
			"func":   sendJSON,
			"action": "json.Marshal",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(j))
}
