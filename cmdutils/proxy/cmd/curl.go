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
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl [OPERATION] [PORT]",
	Short: "curl proxy switcher",
	Run: func(cmd *cobra.Command, args []string) {

	},
	Args: cobra.RangeArgs(1, 2),
}

var curlOnCmd = &cobra.Command{
	Use:   "on",
	Short: "open the curl proxy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			curlOn("1086")
		} else if len(args) == 1 {
			if _, err := strconv.Atoi(args[0]); err != nil {
				fmt.Println("wrong port")
				os.Exit(1)
			}
			curlOn(args[0])
		}
	},
	Args: cobra.RangeArgs(0, 1),
}

var curlOffCmd = &cobra.Command{
	Use:   "off",
	Short: "cancle the proxy",
	Run: func(cmd *cobra.Command, args []string) {
		curlOff()
	},
	Args: cobra.NoArgs,
}

var curlStatusCmd = &cobra.Command{
	Use:   "s",
	Short: "state the curl proxy status",
	Run: func(cmd *cobra.Command, args []string) {
		curlStatus()
	},
	Args: cobra.NoArgs,
}

func curlOn(port string) {
	out, err := os.Create("/Users/ricardo/.curlrc")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	outWriter := bufio.NewWriter(out)
	proxyString := fmt.Sprintf("socks5=\"127.0.0.1:%s\"", port)

	if _, err := outWriter.WriteString(proxyString); err != nil {
		fmt.Println(err)
		return
	}
	outWriter.Flush()
	fmt.Printf("curl on %s", port)
}

func curlOff() {
	err := os.Remove("/Users/ricardo/.curlrc")
	if os.IsNotExist(err) {
		fmt.Println("curl proxy file doesn't exist")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("curl off")
}

func curlStatus() {
	curlrc, err := os.Open("/Users/ricardo/.curlrc")
	if os.IsNotExist(err) {
		fmt.Println("curl proxy file doesn't exist")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}
	defer curlrc.Close()

	curlReader := bufio.NewReader(curlrc)

	curlProfile, err := curlReader.ReadString(byte('\n'))
	// fmt.Print(curlProfile)
	if err != io.EOF && err != nil {
		fmt.Println(err)
		return
	}

	pat := "[0-9]+"
	re, _ := regexp.Compile(pat)
	nums := re.FindAllString(curlProfile, 6)
	fmt.Printf("curl proxy is on %s\n", nums[5])
}

func init() {
	rootCmd.AddCommand(curlCmd)
	curlCmd.AddCommand(curlOnCmd)
	curlCmd.AddCommand(curlOffCmd)
	curlCmd.AddCommand(curlStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// curlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// curlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
