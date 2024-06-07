/*
Copyright © 2024 Focus digitalhydra <digitalhydra@proton.me>
*/
package schedule

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var ScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Show the current schedule",
	Long:  `Root command to create or view the focus schedule `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("schedule called")
	},
}

func init() {
	ScheduleCmd.AddCommand(CreateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
