package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "manage",
		Short: "url shortener at ease",
		Long:  "url shortener application to provide shortened endpoints for provided urls",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateCmd)
}

// sets configuration variables to viper
// always runs on application start up
func initConfig() {

	if cfgFile != "" {
		// use config file provided from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// look for current path for configuration file named  local.env
		viper.SetConfigName("local")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../")

	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("Config file not available, active environment variables will be used instead.")
	} else {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	fmt.Println("--------------------------------------------")
}

