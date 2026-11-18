package config

import platformconfig "github.com/vald3mare/rudstock-platform/config"

type Config struct {
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8083"`
	DBDSN    string `env:"DB_DSN,required,notEmpty"`
	AuthURL  string `env:"AUTH_URL" envDefault:"http://localhost:8082"`
}

func Load() (Config, error) {
	var cfg Config
	err := platformconfig.Load(&cfg)
	return cfg, err
}
