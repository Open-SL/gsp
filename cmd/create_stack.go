package cmd

import (
	"fmt"
	"gsp/internal/commands"

	"github.com/spf13/cobra"
)

var createStackCmd = &cobra.Command{
	Use:   "cs [stack-name] [-d default-branch-name]",
	Short: "Create a new stack",
	Long: `Create a new stack. For example:
	gsp cs my-stack -d default-branch-name`,
	Args: cobra.RangeArgs(1, 2),
	Run:  createStack,
}

func init() {
	rootCmd.AddCommand(createStackCmd)
	createStackCmd.Flags().StringP("defaultBranch", "d", "master", "Default branch name")
}

func createStack(cmd *cobra.Command, args []string) {
	// get stack name from the args
	if len(args) == 0 {
		fmt.Println("Please provide a name for the stack")
		return
	}
	defaultBranch, _ := cmd.Flags().GetString("defaultBranch")
	stackName := args[0]
	fmt.Println("Creating a new stack called", stackName)

	err := commands.CreateStack(stackName, defaultBranch)

	if err != nil {
		fmt.Println("Error writing JSON data:", err)
	}
}
