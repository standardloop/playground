package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

// secret management later
type config struct {
	Port int `env:"APPLICATION_PORT" envDefault:"8080"`

	MySQLEnabled bool   `env:"MYSQL_ENABLED" envDefault:"false"`
	MySQLHost    string `env:"MYSQL_HOST" envDefault:"localhost"`
	MySQLPort    string `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUser    string `env:"MYSQL_USER" envDefault:"root"`
	MySQLPass    string `env:"MYSQL_PASS" envDefault:"mypassword"`
	MySQLDBName  string `env:"MYSQL_DBNAME" envDefault:"playground"`

	PostgresEnabled bool   `env:"POSTGRES_ENABLED" envDefault:"false"`
	PostgresHost    string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort    string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser    string `env:"POSTGRES_USER" envDefault:"root"`
	PostgresPass    string `env:"POSTGRES_PASS" envDefault:"mypassword"`
	PostgresDBName  string `env:"POSTGRES_DBNAME" envDefault:"playground"`
}

func initEnvironment() config {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("initEnvironment() fail")
	}
	return cfg
}

var Env = initEnvironment()
