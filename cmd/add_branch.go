package cmd

import (
	"fmt"
	"gsp/internal/commands"

	"github.com/spf13/cobra"
)

var addBranchCmd = &cobra.Command{
	Use:   "ab [stack-name] [branch-name]",
	Short: "Add a new branch to a stack",
	Long: `Add a new branch to a stack. For example:
	gsp ab my-stack my-branch -o 1`,
	Args: cobra.ExactArgs(2),
	Run:  addBranch,
}

func init() {
	rootCmd.AddCommand(addBranchCmd)
	addBranchCmd.Flags().IntP("order", "o", 0, "Order of the branch")
}

func addBranch(cmd *cobra.Command, args []string) {
	// get stack name and branch name from the args
	stackName := args[0]
	branchName := args[1]
	order, _ := cmd.Flags().GetInt("order")

	err := commands.AddBranch(branchName, stackName, order)

	if err != nil {
		fmt.Println("Error writing JSON data:", err)
	}
}
