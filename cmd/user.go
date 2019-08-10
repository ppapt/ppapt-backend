package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var AdminUserCmd = &cobra.Command{
	Use:   "user",
	Short: "administrate users",
	Long:  `user and its subcommands are used to administrate users directly in the database`,
}

func init() {
	AdminCmd.AddCommand(AdminUserCmd)
}
