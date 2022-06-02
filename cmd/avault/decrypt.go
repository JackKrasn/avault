package main

import (
	"github.com/JackKrasn/avault/cmd/avault/require"
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/spf13/cobra"
)

func newDecryptCmd(cfg *action.Configuration) *cobra.Command {
	dec := action.NewDecrypt(cfg)
	// versionCmd represents the version command
	cmd := &cobra.Command{
		Use:     "decrypt <FILE>",
		Aliases: []string{"dec"},
		Short:   "Decrypt yaml file encrypted by Ansible Vault",
		Long:    `Yaml file encrypted by Ansible Vault. This command decrypts yaml file`,
		Args:    require.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fileName, err := dec.Run(args[0])
			if err != nil {
				return err
			}
			cfg.Log("Decrypted fileName %s", fileName)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVar(&dec.Password, "password", "", "password phrase for decryption")

	return cmd
}
