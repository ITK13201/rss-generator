package domain

import "fmt"

type DatabaseConfig struct {
	User     string `json:"user" env:"DATABASE_USER"`
	Password string `json:"password" env:"DATABASE_PASSWORD"`
	Host     string `json:"host" env:"DATABASE_HOST"`
	Port     string `json:"port" env:"DATABASE_PORT"`
	Database string `json:"database" env:"DATABASE_NAME"`
}

func (config DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}

type Config struct {
	Debug    bool           `json:"debug" env:"DEBUG"`
	Port     string         `json:"port" env:"PORT"`
	Database DatabaseConfig `json:"database"`
}
