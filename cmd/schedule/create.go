/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package schedule

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a schedule to focus on",
	Long:  `Show UI to create the schedule`,
	Run: func(cmd *cobra.Command, args []string) {
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
