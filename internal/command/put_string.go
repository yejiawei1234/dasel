package command

import (
	"github.com/spf13/cobra"
)

func putStringCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "string -f <file> -s <selector> <value>",
		Short: "Update a string property in the given file.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenericPutCommand(genericPutOptions{
				Value:     args[0],
				ValueType: "string",
				Init:      getGenericInit(cmd),
			}, cmd)
		},
	}

	return cmd
}
