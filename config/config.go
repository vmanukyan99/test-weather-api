package config

type Config struct {
	DataBase
}

type DataBase struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	Database string `env:"DB_DATABASE" env-default:"postgres"`
}
