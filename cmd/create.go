package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/zenzora/coveshare/config"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a default config file",
	Run: func(cmd *cobra.Command, args []string) {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = config.GenerateDefaultConfigFile(home + string(os.PathSeparator) + configName + ".yaml")
		if err != nil {
			fmt.Printf("Error creating default config: %s \n", err)
		} else {
			fmt.Println("Default config created")
		}
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
