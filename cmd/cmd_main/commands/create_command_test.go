package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_executeCreateCommand(t *testing.T) {
	createCmd := NewCreateCommandCmd()

	b := bytes.NewBufferString("")
	createCmd.SetOut(b)

	projectName := "test_project"
	dirPath := "/tmp"
	commandName := "ttt"
	createCmd.SetArgs([]string{
		commandName,
	})

	// changing working directory
	changeErr := os.Chdir(filepath.Join(dirPath, projectName))
	if changeErr != nil {
		t.Errorf("Expected to change directory without err, bot got: %v", changeErr)
		return
	}

	createCmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Errorf("Expected to read without err, bot got: %v", err)
		return
	}

	expectedStr := InitializeCreateCommandMessage
	expectedStr += "\n" + fmt.Sprintf(CommandFileCreatedMessage, filepath.Join(dirPath, projectName, "commands", fmt.Sprintf("%s_command.go", commandName)))

	if string(out) != expectedStr {
		t.Errorf("Expected %v, but got: %v", expectedStr, string(out))
		return
	}

	expectedFilePath := filepath.Join(
		dirPath,
		projectName,
		"commands",
		fmt.Sprintf("%s_command.go", commandName),
	)
	if _, err := os.Stat(expectedFilePath); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Filename with the path of `%v` must be existed, but got err: %v", expectedFilePath, err)
		return
	}

	// Read the content of the file

}
