package config

type KafkaConfig struct {
	Brokers       string `toml:"brockers"`
	Topic         string `toml:"topic"`
	Group         string `toml:"group"`
	User          string `toml:"user"`
	Pswd          string `toml:"password"`
	FromBeginning bool   `toml:"from_beginning"`
}
