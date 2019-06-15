package completion

import (
	"github.com/spf13/cobra"
	"os"
)

var commandBash = &cobra.Command{
	Use:   "bash",
	Short: "Output shell completion code for the specified bash shell",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.GenBashCompletion(os.Stdin)
	},
}
