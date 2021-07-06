package config

type Redis struct {
	MaxIdle     int    `mapstructure:"max-idle" json:"max-idle" yaml:"max-idle"`
	MaxActive   int    `mapstructure:"max-active" json:"max-active" yaml:"max-active"`
	IdleTimeout int    `mapstructure:"idle-timeout" json:"idle-timeout" yaml:"idle-timeout"`
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
}
