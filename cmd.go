package cobratools

import (
	"github.com/lylezxl/cobratools/completion"
	"github.com/lylezxl/cobratools/version"
	"github.com/spf13/cobra"
)

func NewCommand(cmd *cobra.Command) {
	cmd.AddCommand(versions.NewCommand())
	cmd.AddCommand(completion.NewCommand())
}
