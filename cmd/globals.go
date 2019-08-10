package cmd

// Port defines the https port to listen on, defaults to 8443.
var Port int

// Filename of the SSL (PEM) Certificate to use for https.
var CertificateFile string

// Filename of the Key (PEM) for the SSL certificate.
var KeyFile string

// Password for the ssl key (currently not supported by golang).
var KeyPassword string

// DatabaseType defines which type of backend database to use: postgres (default), mysql or sqlite3.
var DatabaseType string

// DatabaseName defaults to ppapt.
var DatabaseName string

// DatabaseServer defines on which server the database is located, default is localhost.
var DatabaseServer string

// DatabasePort defines the port foir the database server, defaults to 5432.
var DatabasePort int

// DatabaseUser should be an unprivileged user with select, insert, update, delete rights on the database, default is ppapt.
var DatabaseUser string

// DatabasePassword is the password for the database user (no default).
var DatabasePassword string

// UserEmail is the emeil address for a user
var UserEMail string

// UserName is the name of the User
var UserName string

// UserPassword is the unencrypted password for the user
var UserPassword string

// UserLocked defines if the user can login
var UserLocked bool
