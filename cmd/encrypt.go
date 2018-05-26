package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts objects when called as a helper for the git command",
	Long:  `Encrypts objects when called as a helper for the git command.  You should not call this directly from the command line.  It will be envoked as needed to encrypt files in your git repository when necessary.`,
	Run: func(cmd *cobra.Command, args []string) {
		// wd, err := os.Getwd()
		// if err != nil {
		// 	os.Exit(1)
		// }
		text, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(1)
		}

		log.Printf("Text: %s\n", text)
	},
}
