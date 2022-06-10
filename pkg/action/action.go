package action

type Configuration struct {
	Log func(string, ...interface{})
}

type DebugLog func(format string, v ...interface{})

func (cfg *Configuration) Init(log DebugLog) error {
	cfg.Log = log
	return nil
}
