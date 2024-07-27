package commands

import (
	"gsp/internal/data"
)

func DeleteStack(stackName string) error {
	return data.DeleteJSON(stackName)
}
