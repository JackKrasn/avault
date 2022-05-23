package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cmd, err := newRootCmd(os.Stdout, os.Args[1:])
	if err != nil {
		println("%+v", err)
		os.Exit(1)
	}
	cobra.OnInitialize()
	if err := cmd.Execute(); err != nil {
		println("%+v", err)
	}
}
