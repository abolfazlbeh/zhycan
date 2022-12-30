package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
	"text/template"
	"time"
)

const (
	InitializeMessage         = `Zhycan scaffold project ...`
	RootDirectoryIsCreated    = `Project root (%s) is created ...`
	RootDirectoryIsNotCreated = `Project root (%s) is not created correctly ...%v`
	GoModuleFileIsCreated     = `Go Module File "go.mod" is created ...`
	GoModuleFileIsNotCreated  = `Go Module File "go.mod" is not created ... %v`
	GoModuleIsCreated         = `Go Module "go.mod" is filled ...`
	GoModuleIsNotCreated      = `Go Module "go.mod" is not filled ... %v`
	MainGoFileIsCreated       = `Main program File "main.go" is created ...`
	MainGoFileIsNotCreated    = `Main program File "main.go" is not created ... %v`
	MainGoIsCreated           = `Main program "main.go" is filled ...`
	MainGoIsNotCreated        = `Main program "main.go" is not filled ... %v`
	UserExisted               = "User existed ..."
	UserNotExisted            = "User not existed ... %v"
	SubDirectoryIsNotCreated  = `Sub directory "%s" cannot be created ... %v`
	SubDirectoryIsCreated     = `Sub directory "%s" is created ...`

	DefaultProjectDirectory = "."
)

var ExpectedSubDirectories = func() []string {
	return []string{"controllers", "models", "utils"}
}

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init [application_name]",
		Short: "Create a new bare structure of the application with the name of <application_name>",
		Long:  ``,

		Run: initCmdExecute,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), InitializeMessage)

			// Get project name and directory path
			projectName := args[0]
			projectPath, err := cmd.Flags().GetString("path")
			if err != nil {
				projectPath = DefaultProjectDirectory
			}

			expectedProjectPath := filepath.Join(projectPath, projectName)
			if err := os.Mkdir(expectedProjectPath, os.ModePerm); err != nil {
				fmt.Fprintln(cmd.OutOrStdout())
				fmt.Fprintf(cmd.OutOrStdout(), RootDirectoryIsNotCreated, expectedProjectPath, err)
				return err
			} else {
				fmt.Fprintln(cmd.OutOrStdout())
				fmt.Fprintf(cmd.OutOrStdout(), RootDirectoryIsCreated, expectedProjectPath)
			}

			// Create go.mod file
			goVersion := "1.19"
			err = createGoModFile(cmd, expectedProjectPath, projectName, goVersion)
			if err != nil {
				return err
			}

			// Create main.go file
			currentUser, err := user.Current()
			if err != nil {
				fmt.Fprintln(cmd.OutOrStdout())
				fmt.Fprintf(cmd.OutOrStdout(), UserNotExisted, err)
				return err
			} else {
				fmt.Fprintln(cmd.OutOrStdout())
				fmt.Fprintf(cmd.OutOrStdout(), UserExisted)
			}

			err = createMainGoFile(cmd, expectedProjectPath, projectName, currentUser.Username)
			if err != nil {
				return err
			}

			err = createSubDirectories(cmd, expectedProjectPath)
			if err != nil {
				return err
			}

			return nil
		},
	}
	initCmd.Flags().StringP("path", "p", ".", "The parent path to create a project")
	return initCmd
}

func initCmdExecute(cmd *cobra.Command, args []string) {
}

func createGoModFile(cmd *cobra.Command, expectedProjectPath string, projectName string, goVersion string) error {
	// create go.mod file from template and place it in the directory
	goModPath := filepath.Join(expectedProjectPath, "go.mod")
	file, err := os.Create(goModPath)
	defer file.Close()

	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleFileIsNotCreated, err)
		return err

	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleFileIsCreated)
	}

	temp := template.Must(template.ParseFiles("./templates/gomod.gotmpl"))
	goModuleVars := struct {
		ProjectName string
		Version     string
	}{
		ProjectName: projectName,
		Version:     goVersion,
	}
	err = temp.Execute(file, goModuleVars)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleIsNotCreated, err)
		return err
	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleIsCreated)
	}

	return nil
}

func createMainGoFile(cmd *cobra.Command, expectedProjectPath string, projectName string, creatorUsername string) error {
	// create go.mod file from template and place it in the directory
	mainGoPath := filepath.Join(expectedProjectPath, "main.go")
	file, err := os.Create(mainGoPath)
	defer file.Close()

	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoFileIsNotCreated, err)
		return err

	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoFileIsCreated)
	}

	temp := template.Must(template.ParseFiles("./templates/main.gotmpl"))
	goModuleVars := struct {
		ProjectName     string
		CreatorUserName string
		Time            time.Time
		TimeFormat      string
	}{
		ProjectName:     projectName,
		CreatorUserName: creatorUsername,
		Time:            time.Now().Local(),
		TimeFormat:      time.RFC822,
	}
	err = temp.Execute(file, goModuleVars)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoIsNotCreated, err)
		return err
	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoIsCreated)
	}

	return nil
}

func createSubDirectories(cmd *cobra.Command, expectedProjectPath string) error {
	for _, item := range ExpectedSubDirectories() {
		pathToCreate := filepath.Join(expectedProjectPath, item)

		err := os.Mkdir(pathToCreate, os.ModePerm)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintf(cmd.OutOrStdout(), SubDirectoryIsNotCreated, item, err)
			return err
		} else {
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintf(cmd.OutOrStdout(), SubDirectoryIsCreated, item)
		}
	}

	return nil
}
