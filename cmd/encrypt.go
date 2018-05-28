package cmd

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
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
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
		})

		_, err := ioutil.ReadAll(os.Stdin)
		if err != nil || len(args) == 0 {
			log.Fatal("secretbox: failed to encrypt file on git commit - run `secretbox setup` to ensure that your encryption keys are set up.")
		}

		log.Printf("secretbox: encrypting %s\n", args[0])

	},
}
