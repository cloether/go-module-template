package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-module-template",
	Short: "go-module-template",
	Long:  `Go module template is a demo application microservice.`,
	// Uncomment the following line if your bare application has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	//   cobra.CheckErr(cmd.Help())
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// persistent: [--config, -t]
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "configuration file (default is $HOME/.test-service.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// find the user home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// search config in home directory with name ".test-service" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".test-service")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// if a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if _, err := fmt.Fprintln(os.Stderr, "using config file:", viper.ConfigFileUsed()); err != nil {
			cobra.CheckErr(err)
		}
	} else if errors.As(err, &viper.ConfigFileNotFoundError{}) {
		// Config file not found; ignore error if desired
		log.Println(err)
	} else {
		// cli error
		cobra.CheckErr(err)
	}
}
