package config

type Config struct {
	DataBase
}

type DataBase struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Post     uint16 `env:"DB_PORT" env-default:"6666"`
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_PASSWORD" env-default:"password"`
	Database string `env:"DB_DATABASE" env-default:"db1"`
}
