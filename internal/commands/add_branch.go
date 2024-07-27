package commands

import (
	"fmt"
	"gsp/internal/common"
	"gsp/internal/data"
)

func AddBranch(branchName string, stackName string, order int) error {
	fmt.Println("Adding a new branch called", branchName, "to stack", stackName, "with order", order)

	// Read existing JSON data
	stacks, err := data.ReadJSON()
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return err
	}

	// Check if the stack exists
	stack, exists := stacks[stackName]
	if !exists {
		fmt.Println("Error: Stack does not exist")
		return err
	}

	// Check if the branch already exists
	for _, branch := range stack.Branches {
		if branch.BranchName == branchName {
			fmt.Println("Error: Branch already exists in the stack")
			return err
		}
	}

	// Create a new branch
	newBranch := common.Branch{
		BranchName: branchName,
		Priority:   order,
	}

	// Add the branch to the stack
	stack.Branches = append(stack.Branches, newBranch)

	// Write the updated data to the JSON file
	return data.UpdateJSON(stackName, stack)
}
