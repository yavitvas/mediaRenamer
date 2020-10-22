package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "yaRenamer version",
	Long:  `Show version of the yaRenamer utility.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yaRenamer: version Alpha v.1.2.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
