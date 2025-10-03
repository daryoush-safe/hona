package bootstrap

type Config struct {
	Constants *Constants
	// Add other configuration fields as needed
}

func Run() *Config {
	return &Config{
		Constants: NewConstants(),
	}
}
