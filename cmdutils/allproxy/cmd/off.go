package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(offCmd)
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "turn off proxy",
	Run: func(cmd *cobra.Command, args []string) {
		git := "gitproxy"
		curl := "curlproxy"
		onArgs := []string{"off"}

		if err := exec.Command(git, onArgs...).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := exec.Command(curl, onArgs...).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("All off done!")
	},
}
