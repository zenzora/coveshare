package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zenzora/coveshare/config"
	"github.com/zenzora/coveshare/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		// Lets validate config before starting the server
		validationError := config.Validate()
		if validationError != nil {
			//Config doesn't checkout lets error out
			fmt.Println(validationError)
			os.Exit(1)
		}
		fmt.Println("Serving on port 8080")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
