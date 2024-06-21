package server

type Config struct {
	Host  string `envDefault:""`
	Port  string `envDefault:"8085"`
	Debug bool   `envDefault:"false"`
}
