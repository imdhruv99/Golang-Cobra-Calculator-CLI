package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 != 0 {
				oddSum = oddSum + itemp
			}
		}

		fmt.Printf("The odd addition of %s is %d \n", args, oddSum)
	},
}

func init() {
	RootCmd.AddCommand(oddCmd)
}
