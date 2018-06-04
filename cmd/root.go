package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/secretbox-io/secretbox/config"
	"github.com/secretbox-io/secretbox/providers"
	"github.com/secretbox-io/secretbox/providers/aws"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secretbox",
	Short: "Secretbox is a tool to flexibly and safely manage your secrets",
	Long:  `Secretbox is a tool to flexibly and safely manage your secrets.  More information and documentation is available at https://secretbox.io`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Ran with: %v", args)
	},
}

type cobraCommand func(c *config.Config, p providers.Provider) *cobra.Command

func getProvider(c *config.Config) providers.Provider {
	if c == nil {
		log.Fatal("could not read config file.  Try running `secretbox login` first.")
	}
	if !c.MetaData.IsDefined("provider") {
		log.Fatal("no cloud provider defined.  Try running `secretbox setup`.")
	}
	p, err := aws.NewProvider()
	if err != nil {
		log.Fatal("could not find credentials.  Try running `secretbox setup`.")
	}
	return p
}

func Execute() {

	cfg, _ := config.Read()

	rootCmd.AddCommand(
		encryptCmd(cfg),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
