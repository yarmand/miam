package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "miam",
	Short: "miam is an image manager",
	Long:  `aaa`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "configuration", "c", "", "Configuration file")
	viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("configuration"))
}

// Execute kick the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
