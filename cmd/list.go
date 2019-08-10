/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

// listCmd represents the list command
var AdminUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all users from the database",
	Long:  `Lists all users from the database`,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := Ppapt.ListUsers()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var el int
		var nl int
		for _, line := range u {
			if len(line.EMail) > el {
				el = len(line.EMail)
			}
			if len(line.Name) > nl {
				nl = len(line.Name)
			}
		}
		fs := "%-" + strconv.Itoa(el) + "v %-" + strconv.Itoa(nl) + "v %v\n"
		fmt.Printf(fs, "EMail", "Name", "Locked")
		fmt.Printf(fs, strings.Repeat("-", el), strings.Repeat("-", nl), "-----")
		for _, line := range u {
			fmt.Printf(fs, line.EMail, line.Name, line.Locked)
		}

	},
}

func init() {
	AdminUserCmd.AddCommand(AdminUserListCmd)

}
