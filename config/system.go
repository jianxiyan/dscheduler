package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
}
