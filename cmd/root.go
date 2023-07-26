package cmd

import (
	"github.com/spf13/cobra"
)

const rootCmdLongUsage = `
The Helm KCL Plugin.

* Edit, transformer, validate Helm charts using the helmfile tool and KCL programming language.
`

// New creates a new cobra client
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "helmfile-kcl",
		Short:        "Edit, transformer, validate Helm charts using the helmfile tool and KCL programming language.",
		Long:         rootCmdLongUsage,
		SilenceUsage: true,
	}

	cmd.AddCommand(NewVersionCmd())
	cmd.AddCommand(NewTemplateCmd())
	cmd.SetHelpCommand(&cobra.Command{}) // Disable the help command
	return cmd
}
