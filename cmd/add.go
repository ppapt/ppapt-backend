package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

// AdminUserAddCmd represents the add command for user
var AdminUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user to the database",
	Long: `Add a user with the given email address, name and password and specify if the user is locked.`,
	Run: func(cmd *cobra.Command, args []string) {
		_,err:=Ppapt.NewUser(UserEMail,UserName,UserPassword,UserLocked)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	AdminUserCmd.AddCommand(AdminUserAddCmd)

	AdminUserAddCmd.Flags().StringVar(&UserName, "name", "", "Name for the user")
	AdminUserAddCmd.MarkFlagRequired("name")
	AdminUserAddCmd.Flags().StringVar(&UserPassword, "password", "", "New Password for the user (warning cleartext)")
	AdminUserAddCmd.Flags().BoolVar(&UserLocked, "locked", false, "Should the user be locked")

}
