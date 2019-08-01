// ppapt-backend main.go
//
// ppapt-backend is the rest/webserver for the Pathfinder Pen And Paper Tool
// Using cobra and viper, the main logic is happening within the cmd package
package main

import (
	"github.com/ppapt/ppapt-backend/cmd"
)

func main() {
	cmd.Execute()
}
