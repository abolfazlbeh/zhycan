package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func Test_ExecuteInitCmd(t *testing.T) {
	initCmd := NewInitCmd()

	b := bytes.NewBufferString("")
	initCmd.SetOut(b)

	projectName := "test_project"
	dirPath := "/tmp"
	// check the project folder is created
	expectedPathToCheck := filepath.Join(dirPath, projectName)

	initCmd.SetArgs([]string{
		projectName,
	})
	initCmd.Flags().Set("path", dirPath)
	initCmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Errorf("Expected to read without err, bot got: %v", err)
		return
	}

	expectedStr := InitializeMessage + "\n" + fmt.Sprintf(RootDirectoryIsCreated, expectedPathToCheck)
	expectedStr += "\n" + fmt.Sprintf(GoModuleFileIsCreated)
	expectedStr += "\n" + fmt.Sprintf(GoModuleIsCreated)
	expectedStr += "\n" + fmt.Sprintf(UserExisted)
	expectedStr += "\n" + fmt.Sprintf(MainGoFileIsCreated)
	expectedStr += "\n" + fmt.Sprintf(MainGoIsCreated)

	sortedList := ExpectedSubDirectories()
	sort.Strings(sortedList)

	for _, item := range sortedList {
		expectedStr += "\n" + fmt.Sprintf(SubDirectoryIsCreated, item)
	}
	expectedStr += "\n" + fmt.Sprintf(RootCommandGoFileIsCreated)
	expectedStr += "\n" + fmt.Sprintf(GitInitExecuted)
	expectedStr += "\n" + fmt.Sprintf(GitIgnoreFileIsCreated)

	if string(out) != expectedStr {
		t.Errorf("Expected %v, but got: %v", expectedStr, string(out))
		return
	}

	if _, err := os.Stat(expectedPathToCheck); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Directory with the path of `%v` must be existed, but got err: %v", expectedPathToCheck, err)
		return
	}

	// directory must contain `go.mod` file
	goModePath := filepath.Join(expectedPathToCheck, "go.mod")
	if _, err := os.Stat(goModePath); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Filename with the path of `%v` must be existed, but got err: %v", goModePath, err)
		return
	}

	// directory must contain `main.go` file
	mainGoPath := filepath.Join(expectedPathToCheck, "main.go")
	if _, err := os.Stat(mainGoPath); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Filename with the path of `%v` must be existed, but got err: %v", mainGoPath, err)
		return
	}

	// directory must contain `.git` folder
	dotGitPath := filepath.Join(expectedPathToCheck, ".git")
	st, err := os.Stat(dotGitPath)
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Folder with the path of `%v` must be existed, but got err: %v", dotGitPath, err)
		return
	}
	if err == nil {
		if !st.IsDir() {
			t.Errorf("Folder with the path of `%v` must be existed, but got err: %v", dotGitPath, err)
			return
		}
	}

	// directory must contain `.gitignore` file
	gitIgnorePath := filepath.Join(expectedPathToCheck, ".gitignore")
	if _, err := os.Stat(gitIgnorePath); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Filename with the path of `%v` must be existed, but got err: %v", gitIgnorePath, err)
		return
	}

	// Check these 3 folders exist: controllers, models, configs
	dirs, err := os.ReadDir(expectedPathToCheck)
	if err != nil {
		t.Errorf("Expected to read the sub-directories of %v, but got %v", expectedPathToCheck, err)
	}

	var dirSubsName []string
	for _, item := range dirs {
		if item.IsDir() {
			dirSubsName = append(dirSubsName, item.Name())
		}
	}

	if !reflect.DeepEqual(dirSubsName, sortedList) {
		t.Errorf("The List of Subdirectories --> Expected to be %v, but got %v", sortedList, dirSubsName)
	}
}
