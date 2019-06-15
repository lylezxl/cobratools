package completion

import (
	"github.com/spf13/cobra"
	"os"
)

var commandZsh = &cobra.Command{
	Use:   "zsh",
	Short: "Output shell completion code for the specified zsh shell",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.GenZshCompletion(os.Stdin)
	},
}
