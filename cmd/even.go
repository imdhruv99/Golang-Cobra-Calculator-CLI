package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}

		fmt.Printf("The even addition of %s is %d \n", args, evenSum)
	},
}

func init() {
	RootCmd.AddCommand(evenCmd)
}