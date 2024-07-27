package commands

import (
	"gsp/internal/common"
	"gsp/internal/data"
)

func CreateStack(stackName string, defaultBranch string) error {
	// Initial data to write to the JSON file
	initData := common.Stack{
		StackName:     stackName,
		DefaultBranch: defaultBranch,
	}

	return data.WriteJSON(stackName, initData)
}
