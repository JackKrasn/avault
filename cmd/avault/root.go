package main

import (
	"github.com/spf13/cobra"
	"io"
)

var cfgFile string

var globalUsage = `The utility decrypts yaml files.
Yaml files encrypted by Ansible Vault. 
`

// rootCmd represents the base command when called without any subcommands
//var rootCmd = &cobra.Command{
//	Use:          "avault",
//	Short:        "The utility decrypts yaml files",
//	Long:         globalUsage,
//	SilenceUsage: true,
//}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
//func Execute() {
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "avault",
		Short:        "The utility decrypts yaml files",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	cmd.AddCommand(newDecryptCmd(out))
	cmd.AddCommand(newVersionCmd(out))
	return cmd, nil
}

//func init() {
//	cobra.OnInitialize()
//	rootCmd.AddCommand(newDecryptCmd(os.Stdout))
//	// Cobra also supports local flags, which will only run
//	// when this action is called directly.
//	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//
//}
