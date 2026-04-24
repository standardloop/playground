package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"reflect"
	"strconv"
)

// secret management later
type config struct {
	AppPort  string `json:"appPort" env:"APPLICATION_PORT" envDefault:"8080"`
	LogLevel string `json:"logLevel" env:"LOG_LEVEL" envDefault:"INFO"`

	MySQLEnabled bool   `json:"mySQLEnabled" env:"MYSQL_ENABLED" envDefault:"true"`
	MySQLHost    string `json:"mySQLHost" env:"MYSQL_HOST" envDefault:"localhost"`
	MySQLPort    string `json:"mySQLPort" env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUser    string `json:"mySQLUser" env:"MYSQL_USER" envDefault:"root"`
	MySQLPass    string `json:"mySQLPass" env:"MYSQL_PASS" envDefault:"mypassword"`
	MySQLDBName  string `json:"mySQLDBName" env:"MYSQL_DBNAME" envDefault:"playground"`

	PostgresEnabled bool   `json:"postgresEnabled" env:"POSTGRES_ENABLED" envDefault:"true"`
	PostgresHost    string `json:"postgresHost" env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort    string `json:"postgresPort" env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser    string `json:"postgresUser" env:"POSTGRES_USER" envDefault:"root"`
	PostgresPass    string `json:"postgresPass" env:"POSTGRES_PASS" envDefault:"mypassword"`
	PostgresDBName  string `json:"postgresDBName" env:"POSTGRES_DBNAME" envDefault:"playground"`

	MongoEnabled bool   `json:"mongoEnabled" env:"MONGO_ENABLED" envDefault:"true"`
	MongoHost    string `json:"mongoHost" env:"MONGO_HOST" envDefault:"localhost"`
	MongoPort    string `json:"mongoPort" env:"MONGO_PORT" envDefault:"27017"`
	MongoUser    string `json:"mongoUser" env:"MONGO_USER" envDefault:"root"`
	MongoPass    string `json:"mongoPass" env:"MONGO_PASS" envDefault:"mypassword"`
	MongoDBName  string `json:"mongoDBName" env:"MONGO_DBNAME" envDefault:"playground"`

	RedisEnabled bool   `json:"redisEnabled" env:"REDIS_ENABLED" envDefault:"true"`
	RedisHost    string `json:"redisHost" env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort    string `json:"redisPort" env:"REDIS_PORT" envDefault:"6379"`
	RedisPass    string `json:"redisPass" env:"REDIS_PASS" envDefault:"mypassword"`
	RedisDBNum   string `json:"redisDBNum" env:"REDIS_DB_NUM" envDefault:"0"`

	ShortSha string `json:"shortsha" env:"SHORT_SHA" envDefault:"12345678"`
	GitTag   string `json:"tag" env:"GIT_TAG" envDefault:"v1.0.0"`
}

func parseEnvVars(cfg interface{}) error {
	envVarValue := reflect.ValueOf(cfg).Elem()
	envVarType := envVarValue.Type()

	for i := 0; i < envVarValue.NumField(); i++ {
		field := envVarValue.Field(i)
		fieldType := envVarType.Field(i)

		envKey := fieldType.Tag.Get("env")
		defaultValue := fieldType.Tag.Get("envDefault")

		if envKey == "" {
			continue
		}

		val, exists := os.LookupEnv(envKey)
		if !exists || val == "" {
			val = defaultValue
		}

		if field.CanSet() && field.Kind() == reflect.String {
			field.SetString(val)
		} else if field.CanSet() && field.Kind() == reflect.Bool {
			valAsBool, _ := strconv.ParseBool(val)
			field.SetBool(valAsBool)
		}
	}
	return nil
}

func initEnvironment() config {
	cfg := config{}
	err := parseEnvVars(&cfg)
	if err != nil {
		os.Exit(1)
	}
	return cfg
}

func (conf config) JSON() string {
	jsonData, err := json.Marshal(conf)
	if err != nil {
		slog.Error("Error converting env to json")
		return ""
	}
	return string(jsonData)
}

var Env config = initEnvironment()
