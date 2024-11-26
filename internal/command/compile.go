package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	RunCompileCommandInitMsg  = `Zhycan > Compiling the protobuf file ...`
	RunCompileCommandFileName = `Zhycan > "%s.proto" file is going to be compiled ...`
)

func NewCompileCommandCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "compile [protobuf_name]",
		Short: `Create a protobuf file with the ".proto" extension`,
		Long:  `This command create a directory with the name of protobuf filename, and then compile it to that folder`,
		Run:   runCompileCmdExecute,
		RunE:  runCompileCmdExecuteE,
	}
	return createCmd
}

func runCompileCmdExecuteE(cmd *cobra.Command, args []string) error {
	runCompileCmdExecute(cmd, args)
	return nil
}

func runCompileCmdExecute(cmd *cobra.Command, args []string) {
	fmt.Fprintf(cmd.OutOrStdout(), RunCompileCommandInitMsg+"\n")
	fmt.Fprintf(cmd.OutOrStdout(), RunCompileCommandFileName+"\n", args[0])

	ex, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), err.Error())
	} else {
		fmt.Fprintf(cmd.OutOrStdout(), "Current Path: %s", ex)
	}

}
