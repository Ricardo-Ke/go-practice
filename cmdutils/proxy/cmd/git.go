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
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git proxy switcher",
	Run: func(cmd *cobra.Command, args []string) {

	},
	Args: cobra.RangeArgs(1, 2),
}

var gitOnCmd = &cobra.Command{
	Use:   "on",
	Short: "open the git proxy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			gitOn("1086")
		} else if len(args) == 1 {
			if _, err := strconv.Atoi(args[0]); err != nil {
				fmt.Println("wrong port")
				os.Exit(1)
			}
			gitOn(args[0])
		}
	},
	Args: cobra.RangeArgs(0, 1),
}

var gitOffCmd = &cobra.Command{
	Use:   "off",
	Short: "cancle the git proxy",
	Run: func(cmd *cobra.Command, args []string) {
		gitOff()
	},
	Args: cobra.NoArgs,
}

var gitStatusCmd = &cobra.Command{
	Use:   "s",
	Short: "state the git proxy status",
	Run: func(cmd *cobra.Command, args []string) {
		gitStatus()
	},
	Args: cobra.NoArgs,
}

func gitOn(port string) {
	if err := exec.Command("git", "config", "--global", "http.proxy", "socks5://127.0.0.1:"+port).Run(); err != nil {
		fmt.Println(err)
		return
	}
	if err := exec.Command("git", "config", "--global", "https.proxy", "socks5://127.0.0.1:"+port).Run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("git on %s\n", port)
}

func gitOff() {
	if err := exec.Command("git", "config", "--global", "--unset", "http.proxy").Run(); err != nil {
		fmt.Println(err)
		return
	}
	if err := exec.Command("git", "config", "--global", "--unset", "https.proxy").Run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("git off")
}

func gitStatus() {
	out, err := exec.Command("git", "config", "--get", "http.proxy").CombinedOutput()
	// fmt.Println(string(out))
	if err != nil {
		fmt.Println(err)
	} else if string(out) == "" {
		fmt.Print("git http.proxy is off")
	} else {
		fmt.Printf("git http.proxy is on: %s", string(out))
	}

	out, err = exec.Command("git", "config", "--get", "https.proxy").CombinedOutput()
	// fmt.Println(string(out))
	if err != nil {
		fmt.Println(err)
	} else if string(out) == "" {
		fmt.Print("git https.proxy is off\n")
	} else {
		fmt.Printf("git https.proxy is on: %s", string(out))
	}
}

func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(gitOnCmd)
	gitCmd.AddCommand(gitOffCmd)
	gitCmd.AddCommand(gitStatusCmd)
}
