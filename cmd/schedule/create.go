/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package schedule

import (
	"fmt"
	print "focus/cmd/logger"
	reset "focus/cmd/reset"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type scheduleConfig struct {
	isSchedule bool
	weekDays   [7]string
	startDate  time.Time
	endDate    time.Time
	startTime  time.Time
	endTime    time.Time
	apps       []string
}

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a schedule to focus on",
	Long:  `Show UI to create the schedule`,
	Run: func(cmd *cobra.Command, args []string) {

		print.Logger().Info("Looking for previous schedule....")
		var err error = viper.ReadInConfig()

		if err == nil {
			reset.CreateDefaultConfigFile()
			print.Logger().Info("No config file found", "creating one at", viper.ConfigFileUsed())
		}
		var config scheduleConfig = scheduleConfig{
			isSchedule: viper.GetBool("isschedule"),
			startDate:  viper.GetTime("startdate"),
			endDate:    viper.GetTime("enddate"),
			startTime:  viper.GetTime("starttime"),
			endTime:    viper.GetTime("endtime"),
			weekDays:   [7]string(viper.GetStringSlice("weekdays")),
			apps:       viper.GetStringSlice("apps"),
		}

		cmd.Help()
		fmt.Println("create called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
