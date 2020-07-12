package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println("Serving on port " + viper.GetString("port"))
		server.Serve()
	},
}

func init() {
	serveCmd.Flags().String("base_url", "", "Prefix to format internal links eg (https://my.secret.com or http://localhost:8080)")
	err := viper.BindPFlag("base_url", serveCmd.Flags().Lookup("base_url"))
	if err != nil {
		log.Fatal("Error binding base_url to viper")
	}

	serveCmd.Flags().String("port", "80", "port")
	err = viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	if err != nil {
		log.Fatal("Error binding port to viper")
	}

	rootCmd.AddCommand(serveCmd)
}
