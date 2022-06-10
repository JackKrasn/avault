package cli

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

type EnvSettings struct {
	Debug    bool
	Password string
}

func New() *EnvSettings {
	env := &EnvSettings{
		Password: os.Getenv("AVAULT_PASSWORD"),
	}
	env.Debug, _ = strconv.ParseBool(os.Getenv("AVAULT_DEBUG"))
	return env
}

func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.Debug, "debug", s.Debug, "enable verbose output")
	fs.StringVarP(&s.Password, "password", "p", s.Password, "password phrase for decryption")
}

func (s *EnvSettings) EnvVars() map[string]string {
	envvars := map[string]string{
		"AVAULT_BIN":      os.Args[0],
		"AVAULT_DEBUG":    fmt.Sprint(s.Debug),
		"AVAULT_PASSWORD": s.Password,
	}
	return envvars
}
