package main

import (
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/spf13/cobra"
)

var globalUsage = `The utility decrypts yaml files.
Yaml files are encrypted by Ansible Vault.
`

func newRootCmd(actionConfig *action.Configuration) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "avault",
		Short:        "The utility decrypts yaml files",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	flags := cmd.PersistentFlags()

	settings.AddFlags(flags)
	cmd.AddCommand(newDecryptCmd(actionConfig))
	cmd.AddCommand(newVersionCmd())
	return cmd, nil
}
