package cmd

import (
	"fmt"
	"github.com/ppapt/ppapt-backend/ppapt"
	"github.com/ppapt/ppapt-backend/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// defining the command "server" to start the server component.
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Pen And Paper Tools backend server",
	Long:  `The ppapt backend component`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := ppapt.NewPpapt()
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}
		server.Router(p)
	},
}

// Initialize all the config file parameters and command-line switches.
func init() {
	RootCmd.AddCommand(ServerCmd)
	ServerCmd.Flags().IntVarP(&Port, "port", "p", 8443, "Network port (defaults to 8443)")
	ServerCmd.Flags().StringVarP(&CertificateFile, "certificate-file", "C", "server.crt", "Server certificate file (default server.crt)")
	ServerCmd.Flags().StringVarP(&KeyFile, "key-file", "K", "server.key", "Server certificate keyfile (default server.key)")
	ServerCmd.Flags().StringVarP(&KeyPassword, "key-password", "P", "", "Password for encrypted keyfile (default empty)")
	ServerCmd.Flags().StringVarP(&DatabaseType, "database-type", "t", "postgres", "Type of the database (default postgres)")
	ServerCmd.Flags().StringVarP(&DatabaseName, "database-name", "n", "ppapt", "Name of the database (default ppapt)")
	ServerCmd.Flags().StringVarP(&DatabaseServer, "database-server", "s", "localhost", "Hostname/IP of the database server (default localhost)")
	ServerCmd.Flags().IntVarP(&DatabasePort, "database-port", "o", 5432, "HostnamePort on which the database server listens (default 5432)")
	ServerCmd.Flags().StringVarP(&DatabaseUser, "database-user", "u", "ppapt", "Username for the database connection (default ppapt)")
	ServerCmd.Flags().StringVarP(&DatabasePassword, "database-password", "w", "", "Password for the database connection")

	viper.BindPFlag("port", ServerCmd.Flags().Lookup("port"))
	viper.BindPFlag("certificate-file", ServerCmd.Flags().Lookup("certificate-file"))
	viper.BindPFlag("key-file", ServerCmd.Flags().Lookup("key-file"))
	viper.BindPFlag("database-type", ServerCmd.Flags().Lookup("database-type"))
	viper.BindPFlag("database-name", ServerCmd.Flags().Lookup("database-name"))
	viper.BindPFlag("database-server", ServerCmd.Flags().Lookup("database-server"))
	viper.BindPFlag("database-port", ServerCmd.Flags().Lookup("database-port"))
	viper.BindPFlag("database-user", ServerCmd.Flags().Lookup("database-user"))
	viper.BindPFlag("database-password", ServerCmd.Flags().Lookup("database-password"))

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
