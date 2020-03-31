package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func addInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}

	fmt.Printf("Addition of numbers %s is %d \n", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %f \n", args, sum)
}