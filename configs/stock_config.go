package configs

type StockConfig struct {
	RunWait       int  `mapstructure:"run-wait" yaml:"run-wait" json:"run-wait"`
	InitStock     bool `mapstructure:"init-stock" yaml:"init-stock" json:"init-stock"`
	InitStockCode bool `mapstructure:"init-stock-code" yaml:"init-stock-code" json:"init-stock-code"`
	InitDapan     bool `mapstructure:"init-dapan" yaml:"init-dapan" json:"init-dapan"`
	KeepRun       bool `mapstructure:"keep-run" yaml:"keep-run" json:"keep-run"`
}
