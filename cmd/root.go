package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	// Import the bear package to access BearCmd
	bearCmd "github.com/varunshoor/noteworthy/cmd/bear"
)

var rootCmd = &cobra.Command{
	Use:   "noteworthy",
	Short: "Noteworthy is a Bear note manager",
	Long: `Noteworthy helps you create and manage Forever Notes in Bear.
It provides commands for creating Year, Quarter, Month, and Daily notes.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	
	// Add the Bear command to the root command
	rootCmd.AddCommand(bearCmd.BearCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	// Look for config in home directory
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(fmt.Sprintf("%s/.config/noteworthy", home))
	}
	
	// Read in environment variables
	viper.AutomaticEnv()
	
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}