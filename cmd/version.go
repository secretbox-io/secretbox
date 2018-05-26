package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var CommitHash string
var BuildTime string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of secretbox",
	Long:  `All software has versions. This is secretbox's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nCommit: %s\nBuild Time: %s\n", Version, CommitHash, BuildTime)
	},
}
