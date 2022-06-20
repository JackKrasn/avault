package main

import (
	"fmt"
	"github.com/JackKrasn/avault/pkg/action"
	"github.com/JackKrasn/avault/pkg/cli"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var settings = cli.New()

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}

func warning(format string, v ...interface{}) {
	format = fmt.Sprintf("WARNING: %s\n", format)
	fmt.Fprintf(os.Stderr, format, v...)
}

func main() {
	actionConfig := new(action.Configuration)
	cmd, err := newRootCmd(actionConfig)

	if err != nil {
		warning("%+v", err)
		os.Exit(1)
	}
	cobra.OnInitialize(func() {
		if err := actionConfig.Init(debug); err != nil {
			log.Fatal(err)
		}
	})
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
