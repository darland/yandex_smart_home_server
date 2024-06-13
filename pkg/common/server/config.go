package server

type Config struct {
	Host  string `default:""`
	Port  string `default:"8085"`
	Debug bool   `default:"false"`
}
