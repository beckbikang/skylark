package config

type RedisCfg struct {
	Host           string `toml:"host"`
	Port           int    `toml:"port"`
	ConnectTimeout int    `toml:"connect_timeout"`
	ReadTimeout    int    `toml:"read_timeout"`
	WriteTimeout   int    `toml:"write_timeout"`
	MaxIdle        int    `toml:"max_idle"`
	MaxActive      int    `toml:"max_active"`
	IdleTimeout    int    `toml:"idle_timeout"`
}
