package utility

// ConfigModel конфигурации сервиса
type ConfigModel struct {
	Service struct {
		Version string `json:"version" yaml:"version"`
		Address string `json:"address" yaml:"address"`
		Port    int    `json:"port" yaml:"port"`
	} `json:"service" yaml:"service" toml:"service"`
	Out struct {
		Address string `json:"address" yaml:"address"`
		Port    int    `json:"port" yaml:"port"`
	} `json:"out_service" yaml:"out_service" toml:"out_service"`
}
