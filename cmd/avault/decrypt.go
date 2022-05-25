package main

import (
	"fmt"
	"github.com/JackKrasn/avault/cmd/avault/require"
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/spf13/cobra"
	"io"
)

func newDecryptCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	dec := action.NewDecrypt(cfg)
	// versionCmd represents the version command
	cmd := &cobra.Command{
		Use:     "decrypt <FILE>",
		Aliases: []string{"dec"},
		Short:   "Decrypt yaml file encrypted by Ansible Vault",
		Long:    `Yaml file encrypted by Ansible Vault. This command decrypts yaml file`,
		Args:    require.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.Log("Test")
			fmt.Println("Decryption Data:")
			fmt.Println(args[0])
			fmt.Println(dec.Password)
			if err := dec.Run(args[0]); err != nil {
				return err
			}
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVar(&dec.Password, "password", "", "password phrase for decryption")

	return cmd
}
