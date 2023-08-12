package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	MigrateInitMsg    = `Zhycan > Migrating Sql Models ...`
	MigrateSuccessMsg = `Zhycan > All Sql Models Migrated Successfully...`
	MigrateErrorMsg   = `Zhycan > Migrating Sql Models Encountered Error: %v...`
)

func NewMigrateCmd() *cobra.Command {
	runServerCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrating Sql Models To Pre-Configured Databases...",
		Long:  ``,

		Run:  runServerCmdExecute,
		RunE: runServerCmdExecuteE,
	}
	return runServerCmd
}

func migrateCmdExecuteE(cmd *cobra.Command, args []string) error {
	migrateCmdExecute(cmd, args)
	return nil
}

func migrateCmdExecute(cmd *cobra.Command, args []string) {
	// TODO: in future 'args' must be considered
	fmt.Fprintf(cmd.OutOrStdout(), MigrateInitMsg)

	fmt
}
