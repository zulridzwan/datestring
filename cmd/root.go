/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var dateFormat string
var presetName string
var customDate string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datestring",
	Short: "Usage: datestring [-f] <format> [-p] <preset>",
	Long: `A program to simplify output date value from command line, useful for writing shell scripts.
Usage: datestring [-f] <format> [-p] <preset>`,
	Version: "1.0.0    4 Nov 2021    programmer: zulridzwan@gmail.com",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		//fmt.Println("format:" + dateFormat)
		//fmt.Println("preset:" + presetName)
		if presetName == "today" {
			fmt.Println(now.Format(dateFormat))
		} else {
			var t time.Time

			switch presetName {
			case "yesterday":
				t = now.AddDate(0, 0, -1)
			case "tomorrow":
				t = now.AddDate(0, 0, 1)
			case "monday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now.AddDate(0, 0, -6)
				} else {
					t = now.AddDate(0, 0, -d+1)
				}
			case "friday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now.AddDate(0, 0, -2)
				} else {
					t = now.AddDate(0, 0, -d+5)
				}
			case "saturday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now.AddDate(0, 0, -1)
				} else {
					t = now.AddDate(0, 0, -d+6)
				}
			case "sunday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now
				} else {
					t = now.AddDate(0, 0, -d+7)
				}
			case "next_monday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now.AddDate(0, 0, 1)
				} else {
					t = now.AddDate(0, 0, -d+8)
				}
			case "last_monday":
				d := (int)(now.Weekday())
				if d == 0 {
					t = now.AddDate(0, 0, -6)
				} else {
					t = now.AddDate(0, 0, -d-6)
				}
			case "early_month":
				y, m, _ := now.Date()
				t = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
			case "month_end":
				y, m, _ := now.Date()
				t2 := time.Date(y, m+1, 1, 0, 0, 0, 0, time.Local)
				t = t2.AddDate(0, 0, -1)
			case "next_month":
				y, m, _ := now.Date()
				t = time.Date(y, m+1, 1, 0, 0, 0, 0, time.Local)
			case "last_month":
				y, m, _ := now.Date()
				t = time.Date(y, m-1, 1, 0, 0, 0, 0, time.Local)
			default:
				t = time.Now()
			}
			fmt.Println(t.Format(dateFormat))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.datestring.yaml)")
	//note: default value will be assigned if no flag is specified. Otherwise, you need to pass a value to the flag.
	rootCmd.PersistentFlags().StringVarP(&dateFormat, "format", "f", "2006-01-02", "The time format in Go language specification. https://pkg.go.dev/time#pkg-constants")

	rootCmd.PersistentFlags().StringVarP(&presetName, "preset", "p", "today", `Common relative time name. Available preset:
yesterday
tomorrow
monday
friday
saturday
sunday
next_monday
last_monday
early_month
month_end
next_month
last_month`)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".datestring" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".datestring")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
