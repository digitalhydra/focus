/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Will start, stop, restart the service",
	Long:  `This will run focus on bg and stop the apps and urls and do all the monitoring`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("service called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
