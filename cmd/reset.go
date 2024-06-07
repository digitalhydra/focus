/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package cmd

import (
	"bytes"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset config to default values! CAUTION current schedule will be lost",
	Long:  `for testing or not, use with caution`,
	Run: func(cmd *cobra.Command, args []string) {
		var confirmReset bool
		var configPath string
		configPath = viper.ConfigFileUsed()

		if configPath != "" {
			confirm := huh.NewConfirm().Title("Are you sure you wanna reset: " + configPath).Value(&confirmReset)
			if err := confirm.Run(); err != nil {
				Logger().Error("Confirm Error", "What happen?", err)
				os.Exit(1)
			}

		}
		if confirmReset == true {
			Logger().Warn("Ok, resetting now!")
			createDefaultConfigFile()
		}
	},
}

func createDefaultConfigFile() {

	var defaultConfig = []byte(`
isscheduled: false
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
		Logger().Warn("Set default config options at ", "Location", viper.ConfigFileUsed())
		Logger().Info("Reset done!")
	} else {
		Logger().Error("Reset Error", "Error writing file: "+viper.ConfigFileUsed(), err)
	}

}
func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
