// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all [OPERATION] [PORT]",
	Short: "git and curl proxy switcher",
	Long:  `switch git and curl proxy at the same time`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var allOnCmd = &cobra.Command{
	Use:   "on",
	Short: "open all proxy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			allOn("1086")
		} else if len(args) == 1 {
			if _, err := strconv.Atoi(args[0]); err != nil {
				fmt.Println("wrong port")
				os.Exit(1)
			}
			allOn(args[0])
		}
	},
	Args: cobra.RangeArgs(0, 1),
}

var allOffCmd = &cobra.Command{
	Use:   "off",
	Short: "cancle all proxy",
	Run: func(cmd *cobra.Command, args []string) {
		allOff()
	},
	Args: cobra.NoArgs,
}

var allStatusCmd = &cobra.Command{
	Use:   "s",
	Short: "state all proxy status",
	Run: func(cmd *cobra.Command, args []string) {
		allStatus()
	},
	Args: cobra.NoArgs,
}

func allOn(port string) {
	gitOn(port)
	curlOn(port)
}

func allOff() {
	gitOff()
	curlOff()
}

func allStatus() {
	gitStatus()
	curlStatus()
}

func init() {
	rootCmd.AddCommand(allCmd)

	allCmd.AddCommand(allOnCmd)
	allCmd.AddCommand(allOffCmd)
	allCmd.AddCommand(allStatusCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
