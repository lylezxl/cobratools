package completion

import (
    "github.com/spf13/cobra"
)

var command = &cobra.Command{
    Use:"completion",
    Short:"Output shell completion code for the specified shell (bash or zsh)",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func NewCommand() *cobra.Command  {
    command.AddCommand(commandBash)
    command.AddCommand(commandZsh)
    return command
}