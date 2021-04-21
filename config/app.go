package config

type App struct {
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	Debug           bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Dbtype          string `mapstructure:"Dbtype" json:"Dbtype" yaml:"Dbtype"`
	RuntimeRootPath string `mapstructure:"RuntimeRootPath" json:"RuntimeRootPath" yaml:"RuntimeRootPath"`
	LogSavePath     string `mapstructure:"LogSavePath" json:"LogSavePath" yaml:"LogSavePath"`
	LogSaveName     string `mapstructure:"LogSaveName" json:"LogSaveName" yaml:"LogSaveName"`
	LogFileExt      string `mapstructure:"LogFileExt" json:"LogFileExt" yaml:"LogFileExt"`
	TimeFormat      string `mapstructure:"TimeFormat" json:"TimeFormat" yaml:"TimeFormat"`
}
