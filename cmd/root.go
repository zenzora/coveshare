package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"

  "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)


var cfgFile string
var configName = ".coveshare"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "coveshare",
  Short: "A brief description of your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coveshare.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {


  // Find home directory.
  home, err := homedir.Dir()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {

    viper.AddConfigPath(home)
    viper.AddConfigPath(".")               // optionally look for config in the working directory
    viper.SetConfigName(configName)
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {

    } else {
      fmt.Println(err)
      os.Exit(1)    }
  }
}

