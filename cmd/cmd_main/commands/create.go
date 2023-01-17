package commands

import "github.com/spf13/cobra"

func NewCreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create ",
		Short: `To Creating Different Parts Of The Application`,
		Long:  ``,
	}
	createCmd.AddCommand(NewCreateCommandCmd())
	return createCmd
}
