package cmd

import (
	"fmt"
	"gsp/internal/commands"

	"github.com/spf13/cobra"
)

var deleteStackCmd = &cobra.Command{
	Use:   "ds [stack name]",
	Short: "Delete a stack",
	Long: `Delete a stack from GSP. For example:
	* gsp ds my-stack`,
	Args: cobra.ExactArgs(1),
	Run:  deleteStack,
}

func init() {
	rootCmd.AddCommand(deleteStackCmd)
}

func deleteStack(cmd *cobra.Command, args []string) {
	// get stack name from the args
	stackName := args[0]

	err := commands.DeleteStack(stackName)
	if err != nil {
		fmt.Println("Error deleting stack:", err)
		return
	}

	fmt.Println("Stack deleted successfully")
}
