package cmd

import (
	"fmt"
	"log"
	"os"

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
