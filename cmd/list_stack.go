package cmd

import (
	"fmt"
	"gsp/internal/commands"

	"github.com/spf13/cobra"
)

var listStacksCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all stacks",
	Long: `List all stacks. For example:
	gsp ls`,
	Args: cobra.NoArgs,
	Run:  listStacks,
}

func init() {
	rootCmd.AddCommand(listStacksCmd)
}

func listStacks(cmd *cobra.Command, args []string) {
	err := commands.ListStacks()
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
	}
}
