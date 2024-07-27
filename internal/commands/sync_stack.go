package commands

import (
	"fmt"
	"gsp/internal/cli"
	"gsp/internal/common"
	"os/exec"
	"sort"
)

func SyncStack(stackName string, stack common.Stack, branchName string, defaultBranch string) error {
	sort.Sort(common.ByPriority(stack.Branches))

	// Check if the branch exists
	branchExists, branchIndex := false, -1
	for i, branch := range stack.Branches {
		if branch.BranchName == branchName {
			branchExists, branchIndex = true, i
			break
		}
	}

	if !branchExists && branchName != defaultBranch {
		return fmt.Errorf("branch does not exist in the stack: %s", branchName)
	}

	// Remove all branches before the branch to sync
	filteredBranches := removeElementsBeforeIndex(stack.Branches, branchIndex)

	if branchName == "master" {

		// append master branch to the filteredBranches
		masterBranch := common.Branch{
			BranchName: "master",
			Priority:   -999,
		}
		filteredBranches = append([]common.Branch{masterBranch}, filteredBranches...)
	}

	fmt.Println("Syncing branches:", filteredBranches)
	// Sync the stack
	fmt.Println("Syncing stack", stackName, "against branch", branchName)

	if len(filteredBranches) < 2 {
		return fmt.Errorf("no branches to sync")
	}

	for i := 1; i < len(filteredBranches); i++ {
		err := processBranch(filteredBranches[i].BranchName, filteredBranches[i-1].BranchName)
		if err != nil {
			return fmt.Errorf("error syncing branch %s: %w", filteredBranches[i].BranchName, err)
		}
	}

	return nil
}

func processBranch(branch string, prevBranch string) error {
	// git checkout that branch
	fmt.Printf("checkout branch %s\n", branch)
	if err := cli.RunCommand(exec.Command("git", "checkout", branch)); err != nil {
		return fmt.Errorf("failed to checkout branch %s: %w", branch, err)
	}

	// git pull origin master
	fmt.Printf("pull from origin %s for branch %s\n", prevBranch, branch)
	if err := cli.RunCommand(exec.Command("git", "pull", "--rebase", "origin", prevBranch)); err != nil {
		return fmt.Errorf("failed to pull from origin %s for branch %s: %w", prevBranch, branch, err)
	}

	// git push origin <that-branch> -f
	fmt.Printf("push branch %s\n", branch)
	if err := cli.RunCommand(exec.Command("git", "push", "origin", branch, "-f")); err != nil {
		return fmt.Errorf("failed to push branch %s: %w", branch, err)
	}

	return nil
}

func removeElementsBeforeIndex(items []common.Branch, index int) []common.Branch {
	if index < 0 || index >= len(items) {
		return items // No elements to remove or index out of bounds
	}
	return items[index:]
}
