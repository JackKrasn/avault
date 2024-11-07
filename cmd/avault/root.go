package main

import (
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/spf13/cobra"
)

var globalUsage = `The utility decrypts yaml files.
Yaml files are encrypted by Ansible Vault.

Common actions for Helm:

- avault decrypt:    decrypt yaml file. File is encrypted by Ansible Vault

Environment variables:

| Name                               | Description                                                     |
|------------------------------------|-----------------------------------------------------------------|
| $AVAULT_PASSWORD                   | set a password phrase for decrypting                            |
| $AVAULT_DEBUG                      | enable verbose output                                       	   |
| $AVAULT_DRY                        | replace values with mask instead of decryption                  |
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
