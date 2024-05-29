/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package breaker

import (
	"fmt"

	"github.com/spf13/cobra"
)

// breakCmd represents the break command
var BreakCmd = &cobra.Command{
	Use:   "break",
	Short: "Add a break of 15 or 30 min in the schedule",
	Long:  `The command will create a inBreak falg and a onBreakUntil param in config to bypass the ban until onBreakUntil time`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("break called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breakCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breakCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
