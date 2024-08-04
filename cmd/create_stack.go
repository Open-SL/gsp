package cmd

import (
	"fmt"
	"gsp/internal/commands"

	"github.com/spf13/cobra"
)

var createStackCmd = &cobra.Command{
	Use:   "cs [stack-name] [-p primary-branch-name]",
	Short: "Create a new stack",
	Long: `Create a new stack. For example:
	gsp cs my-stack -p primary-branch-name`,
	Args: cobra.RangeArgs(1, 2),
	Run:  createStack,
}

func init() {
	rootCmd.AddCommand(createStackCmd)
	createStackCmd.Flags().StringP("primaryBranch", "p", "master", "Primary branch name")
}

func createStack(cmd *cobra.Command, args []string) {
	// get stack name from the args
	if len(args) == 0 {
		fmt.Println("Please provide a name for the stack")
		return
	}
	primaryBranch, _ := cmd.Flags().GetString("primaryBranch")
	stackName := args[0]
	fmt.Println("Creating a new stack called", stackName)

	err := commands.CreateStack(stackName, primaryBranch)

	if err != nil {
		fmt.Println("Error writing JSON data:", err)
		return
	}

	fmt.Println("Stack created successfully")
}
