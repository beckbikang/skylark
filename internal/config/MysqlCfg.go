package config

type MysqlCfg struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Pswd         string `toml:"pswd"`
	Db           string `toml:"db"`
	Charset      string `toml:"charset"`
	Locale       string `toml:"locale"`
	Lifetime     int    `toml:"conn_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
}
