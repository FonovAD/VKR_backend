package config

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Port     string         `yaml:"port" required:"true"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" required:"true"`
	Port     int    `yaml:"port" required:"true"`
	Database string `yaml:"database" required:"true"`
	User     string `yaml:"user" required:"true"`
	Password string `yaml:"password"`
}
