package cmd

import (
	"fmt"
	"gsp/internal/commands"
	"gsp/internal/data"

	"github.com/spf13/cobra"
)

var syncStackCmd = &cobra.Command{
	Use:   "ss [stack name] [branch name]",
	Short: "Sync a stack",
	Long: `Sync a stack on remote ✨. For example:
* gsp ss my-stack branch-name
* gsp ss my-stack

If only the stack name is available, all branches in the stack will be synced against master/main.`,
	Args: cobra.RangeArgs(1, 2),
	Run:  syncStack,
}

func init() {
	rootCmd.AddCommand(syncStackCmd)
}

func syncStack(cmd *cobra.Command, args []string) {
	// get stack name from the args
	stackName := args[0]

	// Read existing JSON data
	stacks, err := data.ReadJSON()
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	// Check if the stack exists
	stack, exists := stacks[stackName]
	if !exists {
		fmt.Println("Error: Stack does not exist")
		return
	}

	defaultBranch := stack.DefaultBranch

	// default branch assigned as the default value
	branchName := defaultBranch

	fmt.Println("Syncing stack", stackName, "against branch", branchName)

	if len(args) == 2 && args[1] != "" {
		branchName = args[1]
	}

	err = commands.SyncStack(stackName, stack, branchName, defaultBranch)

	if err != nil {
		fmt.Println("Error syncing stack:", err)
	}
}
