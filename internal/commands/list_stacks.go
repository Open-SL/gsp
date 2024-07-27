package commands

import (
	"fmt"
	"gsp/internal/data"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ListStacks() error {
	stacks, err := data.ReadJSON()
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return err
	}

	if len(stacks) == 0 {
		fmt.Println("No stacks found")
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})

	for _, stack := range stacks {
		table.Append([]string{stack.StackName})
	}
	table.Render() // Send output

	return nil
}
