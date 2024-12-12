package config

type Config struct {
	ENV string `env:"ENV" envDefault:"dev"`
	PORT string ``
}