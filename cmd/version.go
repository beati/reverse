package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version prints the reverse version.",
	Long:  "Version prints the reverse version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("reverse version %s %s\n", version, runtime.Version())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
