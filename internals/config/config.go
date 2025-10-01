package config

type Config struct {
	DSN  string
	Port string
}

func NewConfig() *Config {
	return &Config{
		DSN:  "host=localhost user=cart-service password=pass123 dbname=cart-service-db port=5432 sslmode=disable",
		Port: ":8082",
	}
}
