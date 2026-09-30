package config

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST"`
	Port int    `yaml:"port" env:"SERVER_PORT"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
}
