package whisper

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whisper",
	Short: "whisper - a simple CLI to autogenerate conventional commit messages",
	Long: `whisper is a fast and simple CLI tool for generating commit messages.
You can use whisper to quickly get commit suggestions from your staged git files.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Commit messages")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "oopsiee...there was an error: %s", err)
	}
}
