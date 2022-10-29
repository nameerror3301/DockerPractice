package utility

// ConfigModel конфигурации сервиса
type ConfigModel struct {
	Service struct {
		Version string `json:"version" yaml:"version"`
		Address string `json:"address" yaml:"address"`
		Port    int    `json:"port" yaml:"port"`
	} `json:"service" yaml:"service" toml:"service"`
	Database struct {
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		Database string `json:"database" yaml:"database"`
		Pool     int    `json:"pool" yaml:"pool"`
	} `json:"database" yaml:"database" toml:"database"`
}
