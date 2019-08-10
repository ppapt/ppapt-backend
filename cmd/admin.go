package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// defining the command "admin" to work directly on the backend database.
var AdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Pen And Paper Tools admin tool",
	Long:  `Work directly on the Pen And Paper Tools backend database`,
}

// Initialize all the config file parameters and command-line switches.
func init() {
	RootCmd.AddCommand(AdminCmd)
	AdminCmd.PersistentFlags().StringVarP(&DatabaseType, "database-type", "t", "postgres", "Type of the database (default postgres)")
	AdminCmd.PersistentFlags().StringVarP(&DatabaseName, "database-name", "n", "ppapt", "Name of the database (default ppapt)")
	AdminCmd.PersistentFlags().StringVarP(&DatabaseServer, "database-server", "s", "localhost", "Hostname/IP of the database server (default localhost)")
	AdminCmd.PersistentFlags().IntVarP(&DatabasePort, "database-port", "o", 5432, "HostnamePort on which the database server listens (default 5432)")
	AdminCmd.PersistentFlags().StringVarP(&DatabaseUser, "database-user", "u", "ppapt", "Username for the database connection (default ppapt)")
	AdminCmd.PersistentFlags().StringVarP(&DatabasePassword, "database-password", "w", "", "Password for the database connection")

	viper.BindPFlag("database-type", AdminCmd.PersistentFlags().Lookup("database-type"))
	viper.BindPFlag("database-name", AdminCmd.PersistentFlags().Lookup("database-name"))
	viper.BindPFlag("database-server", AdminCmd.PersistentFlags().Lookup("database-server"))
	viper.BindPFlag("database-port", AdminCmd.PersistentFlags().Lookup("database-port"))
	viper.BindPFlag("database-user", AdminCmd.PersistentFlags().Lookup("database-user"))
	viper.BindPFlag("database-password", AdminCmd.PersistentFlags().Lookup("database-password"))

	viper.SetDefault("database-type", "postgres")
	viper.SetDefault("database-name", "ppapt")
	viper.SetDefault("database-server", "localhost")
	viper.SetDefault("database-port", 5432)
	viper.SetDefault("database-user", "ppapt")
	viper.SetDefault("database-password", "")
}
