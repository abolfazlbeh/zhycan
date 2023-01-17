package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"text/template"
	"time"
)

const (
	InitializeMessage         = `Zhycan > Create project skeleton ...`
	RootDirectoryIsCreated    = `Zhycan > Project root (%s) is created ...`
	RootDirectoryIsNotCreated = `Zhycan > Project root (%s) is not created correctly ...%v`
	GoModuleFileIsCreated     = `Zhycan > Go Module File "go.mod" is created ...`
	GoModuleFileIsNotCreated  = `Zhycan > Go Module File "go.mod" is not created ... %v`
	GoModuleIsCreated         = `Zhycan > Go Module "go.mod" is filled ...`
	GoModuleIsNotCreated      = `Zhycan > Go Module "go.mod" is not filled ... %v`
	MainGoFileIsCreated       = `Zhycan > Main program File "main.go" is created ...`
	MainGoFileIsNotCreated    = `Zhycan > Main program File "main.go" is not created ... %v`
	MainGoIsCreated           = `Zhycan > Main program "main.go" is filled ...`
	MainGoIsNotCreated        = `Zhycan > Main program "main.go" is not filled ... %v`
	UserExisted               = "Zhycan > User existed ..."
	UserNotExisted            = "Zhycan > User not existed ... %v"
	SubDirectoryIsNotCreated  = `Zhycan > Sub directory "%s" cannot be created ... %v`
	SubDirectoryIsCreated     = `Zhycan > Sub directory "%s" is created ...`

	RootCommandGoFileIsCreated    = `Zhycan > Root command File "commands/root.go" is created ...`
	RootCommandGoFileIsNotCreated = `Zhycan > Root command File "commands/root.go" is not created ... %v`
	RootCommandGoIsCreated        = `Zhycan > Root command "commands/root.go" is filled ...`
	RootCommandGoIsNotCreated     = `Zhycan >  Root command "commands/root.go" is not filled ... %v`

	DefaultProjectDirectory = "."
)

var ExpectedSubDirectories = func() []string {
	return []string{"controllers", "models", "utils", "commands"}
}

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init [application_name]",
		Short: "Create a new bare structure of the application with the name of <application_name>",
		Long:  ``,

		Run:  initCmdExecute,
		RunE: initCmdExecuteE,
	}
	initCmd.Flags().StringP("path", "p", ".", "The parent path to create a project")
	return initCmd
}

func initCmdExecuteE(cmd *cobra.Command, args []string) error {
	initCmdExecute(cmd, args)
	return nil
}

func initCmdExecute(cmd *cobra.Command, args []string) {
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
		return
	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), RootDirectoryIsCreated, expectedProjectPath)
	}

	// Create go.mod file
	goVersion := "1.19"
	err = createGoModFile(cmd, expectedProjectPath, projectName, goVersion)
	if err != nil {
		return
	}

	// Create main.go file
	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), UserNotExisted, err)
		return
	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), UserExisted)
	}

	// get current year
	year := time.Now().Year()

	err = createMainGoFile(cmd, expectedProjectPath, projectName, currentUser.Username, year)
	if err != nil {
		return
	}

	err = createSubDirectories(cmd, expectedProjectPath)
	if err != nil {
		return
	}

	err = createRootCommandFile(cmd, expectedProjectPath, projectName, currentUser.Username, year)
	if err != nil {
		return
	}
}

func createGoModFile(cmd *cobra.Command, expectedProjectPath string, projectName string, goVersion string) error {
	// create go.mod file from template and place it in the directory
	goModPath := filepath.Join(expectedProjectPath, "go.mod")
	file, err := os.Create(goModPath)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleFileIsNotCreated, err)
		return err

	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), GoModuleFileIsCreated)
	}
	defer file.Close()

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

func createMainGoFile(cmd *cobra.Command, expectedProjectPath string, projectName string, creatorUsername string, year int) error {
	// create go.mod file from template and place it in the directory
	mainGoPath := filepath.Join(expectedProjectPath, "main.go")
	file, err := os.Create(mainGoPath)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoFileIsNotCreated, err)
		return err

	} else {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), MainGoFileIsCreated)
	}
	defer file.Close()

	temp := template.Must(template.ParseFiles("./templates/main.gotmpl"))
	goModuleVars := struct {
		ProjectName     string
		CreatorUserName string
		Time            time.Time
		TimeFormat      string
		Year            int
	}{
		ProjectName:     projectName,
		CreatorUserName: creatorUsername,
		Time:            time.Now().Local(),
		TimeFormat:      time.RFC822,
		Year:            year,
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
	sortedList := ExpectedSubDirectories()
	sort.Strings(sortedList)

	for _, item := range sortedList {
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

func createRootCommandFile(cmd *cobra.Command, expectedProjectPath string, projectName string, creatorUsername string, year int) error {
	// create go.mod file from template and place it in the directory
	mainGoPath := filepath.Join(expectedProjectPath, "commands", "root.go")
	file, err := os.Create(mainGoPath)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), RootCommandGoFileIsNotCreated, err)
		return err

	}
	defer file.Close()

	temp := template.Must(template.ParseFiles("./templates/root_command.gotmpl"))
	goModuleVars := struct {
		ProjectName     string
		CreatorUserName string
		Time            time.Time
		TimeFormat      string
		Year            int
	}{
		ProjectName:     projectName,
		CreatorUserName: creatorUsername,
		Time:            time.Now().Local(),
		TimeFormat:      time.RFC822,
		Year:            year,
	}
	err = temp.Execute(file, goModuleVars)
	if err != nil {
		fmt.Fprintln(cmd.OutOrStdout())
		fmt.Fprintf(cmd.OutOrStdout(), RootCommandGoIsNotCreated, err)
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout())
	fmt.Fprintf(cmd.OutOrStdout(), RootCommandGoFileIsCreated)

	return nil
}
