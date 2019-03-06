package cmd

import (
	"github.com/ppapt/ppapt-backend/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Port int
var CertificateFile string
var KeyFile string
var KeyPassword string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Pen And Paper Tools backend server",
	Long:  `The ppapt backend component`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Router()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&Port, "port", "P", 9200, "Network port (defaults to 9200)")
	serverCmd.Flags().StringVarP(&CertificateFile, "certificate-file", "C", "server.crt", "Server certificate file (default server.crt)")
	serverCmd.Flags().StringVarP(&KeyFile, "key-file", "K", "server.key", "Server certificate keyfile (default server.key)")
	serverCmd.Flags().StringVarP(&KeyPassword, "key-password", "P", "", "Password for encrypted keyfile (default empty)")

	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("certificate-file", serverCmd.Flags().Lookup("certificate-file"))
	viper.BindPFlag("key-file", serverCmd.Flags().Lookup("key-file"))

	viper.SetDefault("port", 8443)
	viper.SetDefault("certificate-file", "server.crt")
	viper.SetDefault("key-file", "server.key")
	viper.SetDefault("key-password", "")
}
