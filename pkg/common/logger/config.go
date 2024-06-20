package logger

type Config struct {
	Level  string `env:"LEVEL" envDefault:"info"`
	Pretty bool   `env:"PRETTY" envDefault:"false"`
}
