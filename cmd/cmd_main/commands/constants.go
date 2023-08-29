package commands

const (
	DefaultProjectDirectory = "."

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
	AppControllerIsNotCreated = `Zhycan > App "controller.go" cannot be created ... %v`
	AppControllerIsCreated    = `Zhycan > App "controller.go" is created ...`
	AppModelIsNotCreated      = `Zhycan > App "model.go" cannot be created ... %v`
	AppModelIsCreated         = `Zhycan > App "model.go" is created ...`
	AppEngineIsNotCreated     = `Zhycan > App "app.go" cannot be created ... %v`
	AppEngineIsCreated        = `Zhycan > App "app.go" is created ...`

	RootCommandGoFileIsCreated    = `Zhycan > Root command File "commands/root.go" is created ...`
	RootCommandGoFileIsNotCreated = `Zhycan > Root command File "commands/root.go" is not created ... %v`
	RootCommandGoIsNotCreated     = `Zhycan > Root command "commands/root.go" is not filled ... %v`

	GitIgnoreFileIsCreated    = `Zhycan > Git Ignore File ".gitignore" is created ...`
	GitIgnoreFileIsNotCreated = `Zhycan > Git Ignore File ".gitignore" is not created ... %v`
	GitIgnoreIsNotCreated     = `Zhycan > Git Ignore ".gitignore" is not filled ... %v`

	GitInitExecutedError = `Zhycan > Cannot execute git init command ... %v`
	GitInitExecuted      = `Zhycan > Git repository is initialized ...`

	ConfigFileIsCreated       = `Zhycan > Config File "%s" is created ...`
	ConfigFileIsNotCreated    = `Zhycan > Config File "%s" is not created ... %v`
	ConfigDevFileIsCreated    = `Zhycan > Config File "%s" is created for "dev" mode ...`
	ConfigDevFileIsNotCreated = `Zhycan > Config File "%s" is not created for "dev" mode ... %v`

	GoModTidyExecutedError = `Zhycan > Cannot execute go mod tidy command ... %v`
	GoModTidyExecuted      = `Zhycan > "go mod tidy" command is executed ...`

	GreeterProtobufIsNotCreated = `Zhycan > App "greeter.proto" cannot be created ... %v`
	GreeterProtobufIsCreated    = `Zhycan > App "greeter.proto" is created ...`
)

const (
	goModTmpl = `module {{.ProjectName}}

go {{.Version}}

`

	mainTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "main.go" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/
package main

import (
    "fmt"
    "{{.ProjectName}}/commands"
    "github.com/abolfazlbeh/zhycan/pkg/config"
    "github.com/abolfazlbeh/zhycan/pkg/logger"
    "time"
)

/*
The main file of the project
*/
func main() {
    fmt.Println("{{.ProjectName}} is Started ...")

    // config module attributes
    baseConfigPath := "."           // the base path for the parameters. by default it's the current directory
    initialConfigMode := "dev"            // it can be override by environment value --> the value can be "dev" and "prod" and whatever you want
    configPrefix := "{{.ProjectName}}"    // this will be used in reading value from environment with this prefix

    err := config.InitializeManager(baseConfigPath, initialConfigMode, configPrefix)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Testing the logger module works properly
    logger.Log(logger.NewLogObject(
        logger.INFO, "main.go", logger.FuncMaintenanceType, time.Now().UTC(), "Logger Module Works Like A Charm ...", nil))

    // Execute the provided command
    commands.Execute()
}
`

	rootCommandTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "root.go" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/
package commands

import (
	"github.com/spf13/cobra"
	"github.com/abolfazlbeh/zhycan/pkg/cli"
	{{.ProjectName}}/app/app"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
	Short: "A brief description of your application",
	Long: """A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.""",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Use == "runserver" {
			app1 := &app.App{}
			app1.Init()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-corba-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    // Attach Default Zhycan Cli Commands
    cli.AttachCommands(rootCmd)

	// MARK:Commands --- And New Commands Below ---
	// rootCmd.AddCommand(NewInitCmd())
}
`

	gitIgnoreTmpl = `# Create By Zhycan Framework
#
# Copyright © {{.Year}}
# Project: {{.ProjectName}}
# File: ".gitignore" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
# ---------------------------------------------------

### JetBrains template
# Covers JetBrains IDEs: IntelliJ, RubyMine, PhpStorm, AppCode, PyCharm, CLion, Android Studio, WebStorm and Rider
# Reference: https://intellij-support.jetbrains.com/hc/en-us/articles/206544839

# User-specific stuff
.idea/*

# Gradle and Maven with auto-import
# When using Gradle or Maven with auto-import, you should exclude module files,
# since they will be recreated, and may cause churn.  Uncomment if using
# auto-import.
# .idea/artifacts
# .idea/compiler.xml
# .idea/jarRepositories.xml
# .idea/modules.xml
# .idea/*.iml
# .idea/modules
# *.iml
# *.ipr

# CMake
cmake-build-*/

# Mongo Explorer plugin (remove the comment below to include it)
# .idea/**/mongoSettings.xml

# File-based project format
*.iws

# IntelliJ
out/

# mpeltonen/sbt-idea plugin
.idea_modules/

# JIRA plugin
atlassian-ide-plugin.xml

# Cursive Clojure plugin (remove the comment below to include it)
# .idea/replstate.xml

# Crashlytics plugin (for Android Studio and IntelliJ)
com_crashlytics_export_strings.xml
crashlytics.properties
crashlytics-build.properties
fabric.properties

# Editor-based Rest Client (remove the comment below to include it)
# .idea/httpRequests

# Android studio 3.1+ serialized cache file (remove the comment below to include it)
# .idea/caches/build_file_checksums.ser

### Go template
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

### VisualStudioCode template
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
*.code-workspace

# Local History for Visual Studio Code
.history/

### Xcode template
# Xcode
#
# gitignore contributors: remember to update Global/Xcode.gitignore, Objective-C.gitignore & Swift.gitignore

## User settings
xcuserdata/

## compatibility with Xcode 8 and earlier (ignoring not required starting Xcode 9)
*.xcscmblueprint
*.xccheckout

## compatibility with Xcode 3 and earlier (ignoring not required starting Xcode 4)
build/
DerivedData/
*.moved-aside
*.pbxuser
!default.pbxuser
*.mode1v3
!default.mode1v3
*.mode2v3
!default.mode2v3
*.perspectivev3
!default.perspectivev3

## Gcc Patch
/*.gcno
`

	baseConfigTmpl = `{
  "name": "{{.ProjectName}}",
  "config_must_watched": true,
  "config_remote_addr": "0.0.0.0:7777",
  "config_remote_infra": "grpc",
  "config_remote_duration": 300,
  "modules": [
    {"name":"logger", "type": "local"}
  ]
}`
	loggerConfigTmpl = `{
  "type": "zap",
  "outputs": ["console", "file"],
  "channel_size": 1000,
  "options": ["caller", "stackTrace"],
  "console": {
    "level": "debug"
  },
  "file": {
    "level": "debug",
    "path": "/tmp"
  },
  "graylog": {
    "ip": "graylog.graylog",
    "port": 12201,
    "stdout": true
  },
  "syslog": {
    "ip": "172.25.205.37",
    "port": 514,
    "ctype": "tcp"
  }
}`
	httpConfigTmpl = `{
  "default": "s1",
  "servers": [
    {
      "name":                   "s1",
      "addr":                   ":3000",
      "versions":               ["v1", "v2"],
      "conf": {
        "server_header": "",
        "strict_routing": false,
        "case_sensitive": false,
        "unescape_path": false,
        "etag": false,
        "body_limit": 4194304,
        "concurrency": 262144,
        "read_timeout": -1,
        "write_timeout": -1,
        "idle_timeout": -1,
        "read_buffer_size": 4096,
        "write_buffer_size": 4096,
        "compressed_file_suffix": ".gz",
        "get_only": false,
        "disable_keepalive": false,
        "network": "tcp",
        "enable_print_routes": true,
        "attach_error_handler": true
      }
    }
  ]
}`

	dbConfigTmpl = `{
  "connections": ["server1"],
  "server1": {
    "type": "sqlite",
    "db": "file.db",
    "options": {
      "_fk": "1"
    },
    "config": {
      "skip_default_transaction": false,
      "dry_run": false,
      "prepare_stmt": true,
      "disable_automatic_ping": false,
      "disable_foreign_key_constraint_when_migrating": false,
      "ignore_relationships_when_migrating": false,
      "disable_nested_transaction": false
    },
    "logger": {
      "slow_threshold": 1000,
      "ignore_record_not_found_error": false,
      "parameterized_queries": false,
      "log_level": "debug"
    }
  },
  "server2": {
    "type": "mysql",
    "db": "databasename",
    "username": "username",
    "password": "password",
    "host": "127.0.0.1",
    "port": "3306",
    "protocol": "tcp",
    "options": {
      "charset": "utf8mb4",
      "parseTime": "True",
      "loc": "Local"
    },
    "specific_config": {
      "default_string_size": 256,
      "disable_datetime_precision": true,
      "support_rename_index": true,
      "support_rename_column": true,
      "skip_initialize_with_version": false,
      "disable_with_returning": false,
      "support_for_share_clause": false,
      "support_null_as_default_value": false,
      "support_rename_column_unique": false,
      "default_datetime_precision": 3
    }
  },
  "server3": {
    "type": "postgresql",
    "db": "option_trading",
    "username": "username",
    "password": "password",
    "host": "127.0.0.1",
    "port": "3306",
    "options": {
      "sslmode": "disable",
      "TimeZone": "Asia/Tehran"
    },
    "specific_config": {
      "prefer_simple_protocol": true,
      "without_returning": false
    }
  },
  "server4": {
    "type": "mongodb",
    "db": "",
    "username": "username",
    "password": "password",
    "host": "127.0.0.1",
    "port": "27017",
    "options": {
      "maxPoolSize": "100",
      "w": "majority",
      "connectTimeoutMS": "30000"
    },
    "logger": {
      "component_command": "debug",
      "component_connection": "info",
      "max_document_length": 1000
    }
  }
}`

	appControllerTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "app/controller.go" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/

package app

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abolfazlbeh/zhycan/pkg/http"
)

// MARK: Controller

// SampleController - a sample controller to show the functionality
type SampleController struct {}

// GetName - Get name of the controller, you can return empty string if you want
func (s *SampleController) GetName() string {
	return "Sample"
}

// Routes - returning controller specific routes to be registered
func (ctrl *SampleController) Routes() []http.HttpRoute {
    return []http.HttpRoute {
        http.HttpRoute {
            Method:     http.MethodGet,
            Path:       "/hello",
            RouteName:  "hello",
            F:          &ctrl.GetHello,
        },
    }
}

// GetHello - just return the 'Hello World' string to user
func (ctrl *SampleController) GetHello(c *fiber.Ctx) error {
    return c.SendString("Hello World")
}

`

	appEngineTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "app/app.go" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/

package app

import (
    "github.com/abolfazlbeh/zhycan/pkg/engine"
    "github.com/abolfazlbeh/zhycan/pkg/http"
)

// MARK: App Engine

// App - application engine structure that must satisfy one of the engine interface such as 'engine.RestfulApp', ...
type App struct {}

// Init - initialize the app
func (app *App) Init() {
    err := engine.RegisterRestfulController(&SampleController{})
    if err != nil {
        logger.Log(logger.NewLogObject(
            logger.ERROR, "App.Init", logger.FuncMaintenanceType, time.Now().UTC(), "Cannot Register Restful Controller", err))
    }

    http.PrintAllRoutes()
}
`

	appModelTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "app/model.go" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/

package app

import (
    "github.com/gofiber/fiber/v2"
    "github.com/abolfazlbeh/zhycan/pkg/db"
)

// MARK: Models

// User - a sample model to show the functionality
type User struct {
    gorm.Model
    Name string
}

// CreateNewUser - create a new user record in database
func CreateNewUser(name string) (*User, int64, error) {
    u := User{Name: "test"}

    database, err := db.GetDb("default")
    if err != nil {
        return nil, 0, errors.New("UserCreateError")
    }

    result := database.Create(&u)
    if result.Error != nil {
        return nil, 0, result.Error
    }

    return &u, result.RowsAffected, nil
}

// GetAllUsers - get all user records from database
func GetAllUsers() (*[]User, int64, error) {
    database, err := db.GetDb("default")
    if err != nil {
        return nil, errors.New("UserCreateError")
    }

    var users []User

    result := db2.Find(&users)
    if result.Error != nil {
        return nil, 0, result.Error
    }

    return &users, result.RowsAffected, nil
}
`

	greeterProtobufTmpl = `/*
Create By Zhycan Framework

Copyright © {{.Year}}
Project: {{.ProjectName}}
File: "app/proto/greeter.proto" --> {{ .Time.Format .TimeFormat }} by {{.CreatorUserName}}
------------------------------
*/

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloResponse {
  string message = 1;
}
`
)
