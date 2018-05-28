package cmd

import (
	"io/ioutil"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

type encryptStatus struct {
	numFiles int
	sync.Once
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts objects when called as a helper for the git command",
	Long:  `Encrypts objects when called as a helper for the git command.  You should not call this directly from the command line.  It will be envoked as needed to encrypt files in your git repository when necessary.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)
		text, err := ioutil.ReadAll(os.Stdin)
		if err != nil || len(args) == 0 {
			log.Fatal("secretbox: failed to encrypt file on git commit - run `secretbox setup` to ensure that your encryption keys are set up.")
		}
		log.WithFields(log.Fields{
			"command": "encrypt",
			"file":    args[0],
		}).Debugf("starting encryption")

		log.Printf("secretbox: encrypting %s\n", args[0])
	},
}
