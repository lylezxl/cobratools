package versions

import (
    "github.com/spf13/cobra"
    "fmt"
)

var command = &cobra.Command{
    Use:"version",
    Short:"version info",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("%s %#v\n",project,getVersionInfo())
    },
}

func NewCommand() *cobra.Command  {
    return command
}