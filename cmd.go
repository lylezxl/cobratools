package cobratools

import (
    "github.com/spf13/cobra"
    "github.com/lylezxl/cobratools/version"
    "github.com/lylezxl/cobratools/completion"
)
func NewCommand(cmd *cobra.Command)  {
    cmd.AddCommand(versions.NewCommand())
    cmd.AddCommand(completion.NewCommand())
}
