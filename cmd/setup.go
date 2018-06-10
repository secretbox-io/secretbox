package cmd

import (
	"github.com/secretbox-io/secretbox/config"
	"github.com/spf13/cobra"
)

func setupCmd() *cobra.Command {

	c := config.MustRead()
	if !c.MetaData.IsDefined("provider") {
		c.Provider = config.AWS
		c.Write()
	}

	p := getProvider(c)

	return &cobra.Command{
		Use:   "setup",
		Short: "Set up your cloud provider to work with secretbox",
		Long:  `Set up your cloud provider to work with secretbox.  Creates customer master keys and roles to enable your production instances to decrypt secrets on demand.`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
