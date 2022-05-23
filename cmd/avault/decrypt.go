package main

import (
	"fmt"
	"github.com/JackKrasn/avault/cmd/avault/require"
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/spf13/cobra"
	"io"
)

func newDecryptCmd(out io.Writer) *cobra.Command {
	dec := action.NewDecrypt()
	// versionCmd represents the version command
	cmd := &cobra.Command{
		Use:     "decrypt <FILE>",
		Aliases: []string{"dec"},
		Short:   "Decrypt yaml file encrypted by Ansible Vault",
		Long:    `Yaml file encrypted by Ansible Vault. This command decrypts yaml file`,
		Args:    require.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Decryption Data:")
		},
	}

	f := cmd.Flags()
	f.StringVar(&dec.Password, "password", "", "password phrase for decryption")

	return cmd
}
