package configs

type RedisConfig struct {
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr"`
	DB   int    `mapstructure:"db" yaml:"db" json:"db"`
	Pwd  string `mapstructure:"password" yaml:"password" json:"password"`
}
