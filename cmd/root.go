package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ngmagic",
	Short: "ngMagic es una fiesta de cli",
	Long:  `por ahora solo agrega proxy_pass a nginx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ngMagic! run ngMagic --help for details")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
