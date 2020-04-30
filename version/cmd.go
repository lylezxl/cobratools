package versions

import (
	"fmt"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %#v\n", App, getVersionInfo())
	},
}

func NewCommand() *cobra.Command {
	return command
}
