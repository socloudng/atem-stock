package configs

type MongoConfig struct {
	Server string `mapstructure:"server" yaml:"server" json:"server"`
	Port   string `mapstructure:"port" yaml:"port" json:"port"`
	User   string `mapstructure:"user" yaml:"user" json:"user"`
	Pwd    string `mapstructure:"password" yaml:"password" json:"password"`
}
