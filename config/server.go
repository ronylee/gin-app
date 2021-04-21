package config

type Server struct {
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	Pprof   string `mapstructure:"pprof" json:"pprof" yaml:"pprof"`
}
