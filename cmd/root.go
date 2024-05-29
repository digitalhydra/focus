/*
Copyright © 2024 focus
*/
package cmd

import (
	"bytes"
	"fmt"
	breaker "focus/cmd/break"
	exception "focus/cmd/exception"
	schedule "focus/cmd/schedule"
	service "focus/cmd/service"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "focus",
	Short: "A tool to set a schedule to focus, a CLI tool for ADHD",
	Long:  `Once a schedule is setup a service will stop apps and urls from running so you can focus`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	addSubcommandPalletes()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.focus.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Print("initconfig ")
	if cfgFile != "" {
		fmt.Printf("cfgfile!=empty")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Printf(" else setup viper")
		// Find home directory.
		configPath, err := os.UserConfigDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".focus" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".focus")

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf(" else config")
		createDefaultConfigFile()
	}
}

func addSubcommandPalletes() {
	rootCmd.AddCommand(schedule.ScheduleCmd)
	rootCmd.AddCommand(service.ServiceCmd)
	rootCmd.AddCommand(exception.ExceptionCmd)
	rootCmd.AddCommand(breaker.BreakCmd)
}

func createDefaultConfigFile() {

	var defaultConfig = []byte(`
startdate: 01/01/2024
enddate: null
weekdays:
- monday
- tuesday
- wednesday
- thursday
- friday
apps:
- leather
- denim
starttime: 8:00am
endtime: 4:00pm
onbreak: false
exception: 
- 15/05/2024
`)
	viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	err := viper.SafeWriteConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "set default config file:", viper.ConfigFileUsed())
	}

}
