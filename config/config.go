package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV string `env:"ENV" envDefault:"dev"`
	PORT string `env:"PORT" envDefault:"8080"`
	PgConfig PGConfig `envPrefix:"POSTGRES_"`
	JWTConfig JWTConfig `envPrefix:"JWT_"`
	// MysqlConfig MySQLConfig `envPrefix:"MYSQL_"`
	// SMTPConfig SMTPConfig `envPrefix:"SMTP_"`
}

type JWTConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

type PGConfig struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"5432"`
	User string `env:"USER" envDefault:"sumagesic"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
	SSLMode string `env:"SSLMODE" envDefault:"disable"`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}