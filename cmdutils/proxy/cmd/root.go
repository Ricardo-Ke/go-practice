package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defaultPort = 1086

var rootCmd = &cobra.Command{
	Use:   "proxy [OPERATION] [PORT]",
	Short: "git, curl and wget 代理切换工具",
}

// Execute is the main entry
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
