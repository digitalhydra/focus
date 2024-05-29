/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package exception

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exceptionCmd represents the exception command
var ExceptionCmd = &cobra.Command{
	Use:   "exception",
	Short: "Will add a exception for the day in a schedule in case of holidays",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("exception called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exceptionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exceptionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
