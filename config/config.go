package config

// Config is a application configuration structure
type Config struct {
	Database struct {
		Host       string `env:"DB_HOST" env-description:"Database host"`
		Port       string `env:"DB_PORT" env-description:"Database port"`
		Username   string `env:"DB_USER" env-description:"Database user name"`
		Password   string `env:"DB_PASSWORD" env-description:"Database user password"`
		Name       string `env:"DB_NAME" env-description:"Database name"`
		Connection string `env:"DB_CONNECTIONS"`
	}
}
