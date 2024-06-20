package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	cfg, err := env.ParseAsWithOptions[Config](env.Options{Prefix: "LOG_"})
	if err != nil {
		panic(fmt.Sprintf("failed to parse config: %v", err))
	}

	fmt.Printf("%+v\n", cfg)

	logLevel, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		panic(fmt.Sprintf("failed to parse log level: %v", err))
	}

	var output io.Writer

	output = os.Stdout

	if cfg.Pretty {
		output = zerolog.NewConsoleWriter()
	}

	return zerolog.New(output).Level(logLevel).With().Timestamp().Logger()
}
