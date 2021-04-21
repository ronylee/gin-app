package config

type Base struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
}
