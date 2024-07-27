package data

import (
	"encoding/json"
	"fmt"
	"gsp/internal/common"
	"gsp/internal/constants"
	"os"
	"path/filepath"
)

type Config = map[string]common.Stack

func WriteJSON(stackName string, data common.Stack) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	// Define the .gsp directory path inside the home directory
	gspDir := filepath.Join(homeDir, constants.FolderPath)

	// Create the .gsp directory if it doesn't exist
	err = os.MkdirAll(gspDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating .gsp directory:", err)
		return err
	}

	// Define the file path to data.json inside the .gsp directory
	filePath := filepath.Join(gspDir, constants.FileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file does not exist, create it with initial data
		initialData := Config{
			stackName: data,
		}
		// Marshal initial data to JSON
		fileData, err := json.MarshalIndent(initialData, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling initial data:", err)
			return err
		}

		// Ensure the directory exists
		dir := filepath.Dir(filePath)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		// Write initial data to the file
		return os.WriteFile(filePath, fileData, 0644)
	}

	// Read the existing JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Unmarshal JSON data into Config struct
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return err
	}

	// Check if the key already exists
	if _, exists := config[stackName]; exists {
		fmt.Println("Error: Key already exists in the file.")
		return err
	}

	// Add the key-value pair to the Config
	config[stackName] = data

	// Marshal updated data to JSON
	updatedFileData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling updated data:", err)
		return err
	}

	// Write the updated data back to the file
	return os.WriteFile(filePath, updatedFileData, 0644)
}

func ReadJSON() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return nil, err
	}

	// Define the .gsp directory path inside the home directory
	gspDir := filepath.Join(homeDir, constants.FolderPath)

	// Create the .gsp directory if it doesn't exist
	err = os.MkdirAll(gspDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating .gsp directory:", err)
		return nil, err
	}

	// Define the file path to data.json inside the .gsp directory
	filePath := filepath.Join(gspDir, constants.FileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist.")
		return nil, err
	}

	// Read the existing JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Unmarshal JSON data into Config struct
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return nil, err
	}

	return config, nil
}

func UpdateJSON(stackName string, data common.Stack) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	// Define the .gsp directory path inside the home directory
	gspDir := filepath.Join(homeDir, constants.FolderPath)

	// Create the .gsp directory if it doesn't exist
	err = os.MkdirAll(gspDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating .gsp directory:", err)
		return err
	}

	// Define the file path to data.json inside the .gsp directory
	filePath := filepath.Join(gspDir, constants.FileName)

	// Read the existing JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Unmarshal JSON data into Config struct
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return err
	}

	// Add the key-value pair to the Config
	config[stackName] = data

	// Marshal updated data to JSON
	updatedFileData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling updated data:", err)
		return err
	}

	// Write the updated data back to the file
	return os.WriteFile(filePath, updatedFileData, 0644)
}

func DeleteJSON(stackName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	// Define the .gsp directory path inside the home directory
	gspDir := filepath.Join(homeDir, constants.FolderPath)

	// Create the .gsp directory if it doesn't exist
	err = os.MkdirAll(gspDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating .gsp directory:", err)
		return err
	}

	// Define the file path to data.json inside the .gsp directory
	filePath := filepath.Join(gspDir, constants.FileName)

	// Read the existing JSON file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Unmarshal JSON data into Config struct
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return err
	}

	// delete stack from the Config
	delete(config, stackName)

	// Marshal updated data to JSON
	updatedFileData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling updated data:", err)
		return err
	}

	// Write the updated data back to the file
	return os.WriteFile(filePath, updatedFileData, 0644)
}
