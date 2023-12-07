package main

import (
	"fmt"

	"github.com/JackKrasn/avault/internal/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	var short bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of generated code example",
		Long:  `All software has versions. This is generated code example`,
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Println(version.Version)
				return
			}
			fmt.Println("Build Date:", version.BuildDate)
			fmt.Println("Git Commit:", version.GitCommit)
			fmt.Println("Version:", version.Version)
			fmt.Println("Go Version:", version.GoVersion)
			fmt.Println("OS / Arch:", version.OsArch)
		},
	}

	f := cmd.Flags()
	f.BoolVar(&short, "short", false, "print the version number")

	return cmd
}
