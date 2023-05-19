/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// customCmd represents the custom command
var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Usage: datestring [-f] <format> custom [-y] <int> [-m] <int> [-d] <int>",
	Long: `Generates a custom date based on the values of year, month and day relative to the current date. 
Usage: datestring [-f] <format> custom [-y] <int> [-m] <int> [-d] <int>

Example:
Return a date for 3 days from now: datestring custom -d 3
Return a date for last month: datestring custom -m -1
Return a date 1 year, 1 month and 1 day from now: datestring custom -y 1 -m 1 -d 1`,
	Run: func(cmd *cobra.Command, args []string) {
		var no_args bool
		now := time.Now()
		yy, e := cmd.Flags().GetInt("year")
		mm, e := cmd.Flags().GetInt("month")
		dd, e := cmd.Flags().GetInt("day")

		if e != nil {
			fmt.Println(now.Format(dateFormat))
		} else {
			no_args = yy == 0 && mm == 0 && dd == 0
		}

		if no_args {
			fmt.Println(now.Format(dateFormat))
		} else {
			t := now.AddDate(yy, mm, dd)
			fmt.Println(t.Format(dateFormat))
		}

	},
}

func init() {
	rootCmd.AddCommand(customCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// customCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	customCmd.Flags().IntP("year", "y", 0, "Add year to the current date. Use negative value for past year.")
	customCmd.Flags().IntP("month", "m", 0, "Add month to the current date. Use negative value for past month.")
	customCmd.Flags().IntP("day", "d", 0, "Add day to the current date. Use negative value for past day.")
}
