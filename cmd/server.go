package cmd

import (
	"github.com/ppapt/ppapt-backend/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

// defining the command "server" toi start the server component.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Pen And Paper Tools backend server",
	Long:  `The ppapt backend component`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Router()
	},
}

// Initialize all the config file parameters and command-line switches.
func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&Port, "port", "p", 8443, "Network port (defaults to 8443)")
	serverCmd.Flags().StringVarP(&CertificateFile, "certificate-file", "C", "server.crt", "Server certificate file (default server.crt)")
	serverCmd.Flags().StringVarP(&KeyFile, "key-file", "K", "server.key", "Server certificate keyfile (default server.key)")
	serverCmd.Flags().StringVarP(&KeyPassword, "key-password", "P", "", "Password for encrypted keyfile (default empty)")
	serverCmd.Flags().StringVarP(&DatabaseType, "database-type", "t", "postgres", "Type of the database (default postgres)")
	serverCmd.Flags().StringVarP(&DatabaseName, "database-name", "n", "ppapt", "Name of the database (default ppapt)")
	serverCmd.Flags().StringVarP(&DatabaseServer, "database-server", "s", "localhost", "Hostname/IP of the database server (default localhost)")
	serverCmd.Flags().IntVarP(&DatabasePort, "database-port", "o", 5432, "HostnamePort on which the database server listens (default 5432)")
	serverCmd.Flags().StringVarP(&DatabaseUser, "database-user", "u", "ppapt", "Username for the database connection (default ppapt)")
	serverCmd.Flags().StringVarP(&DatabasePassword, "database-password", "w", "", "Password for the database connection")

	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("certificate-file", serverCmd.Flags().Lookup("certificate-file"))
	viper.BindPFlag("key-file", serverCmd.Flags().Lookup("key-file"))
	viper.BindPFlag("database-type", serverCmd.Flags().Lookup("database-type"))
	viper.BindPFlag("database-name", serverCmd.Flags().Lookup("database-name"))
	viper.BindPFlag("database-server", serverCmd.Flags().Lookup("database-server"))
	viper.BindPFlag("database-port", serverCmd.Flags().Lookup("database-port"))
	viper.BindPFlag("database-user", serverCmd.Flags().Lookup("database-user"))
	viper.BindPFlag("database-password", serverCmd.Flags().Lookup("database-password"))

	viper.SetDefault("port", 8443)
	viper.SetDefault("certificate-file", "server.crt")
	viper.SetDefault("key-file", "server.key")
	viper.SetDefault("key-password", "")
	viper.SetDefault("database-type", "postgres")
	viper.SetDefault("database-name", "ppapt")
	viper.SetDefault("database-server", "localhost")
	viper.SetDefault("database-port", 5432)
	viper.SetDefault("database-user", "ppapt")
	viper.SetDefault("database-password", "")
}
