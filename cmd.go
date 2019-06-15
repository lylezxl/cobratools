package cobra-tools

import (
    "github.com/spf13/cobra"
    "github.com/lylezxl/cobra-tools/version"
    "github.com/lylezxl/cobra-tools/completion"
)
func NewCommand(cmd *cobra.Command)  {
    cmd.AddCommand(versions.NewCommand())
    cmd.AddCommand(completion.NewCommand())
}
