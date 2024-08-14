package configs

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	runMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config {
	cfgPath := GetConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := ParsConfig(v)
	if err != nil {
		log.Fatal("Err on parse config %v", err)
	}
	return cfg
}

func ParsConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, err
}

func LoadConfig(filename string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}
	return v, nil
}

func GetConfigPath(env string) string {
	if env == "docker" {
		return "configs/config-docker"
	} else if env == "production" {
		return "configs/config-production"
	} else {
		return "/src/configs/config-development"
	}
}
