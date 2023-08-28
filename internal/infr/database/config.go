package database

type PostgresConfig struct {
	Server       string `envconfig:"SERVER"`
	Port         uint16 `envconfig:"PORT"`
	User         string `envconfig:"USER"`
	Password     string `envconfig:"PASSWORD"`
	DatabaseName string `envconfig:"DB_NAME"`
}
