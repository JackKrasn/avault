package cli

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

type EnvSettings struct {
	Debug bool
}

func New() *EnvSettings {
	env := &EnvSettings{}
	env.Debug, _ = strconv.ParseBool(os.Getenv("AVAULT_DEBUG"))

	return env
}

func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.Debug, "debug", s.Debug, "enable verbose output")
}

func (s *EnvSettings) EnvVars() map[string]string {
	envvars := map[string]string{
		"AVAULT_BIN":   os.Args[0],
		"AVAULT_DEBUG": fmt.Sprint(s.Debug),
	}
	return envvars
}
