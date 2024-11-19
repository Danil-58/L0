package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App   App   
	HTTP  HTTP  
	DB    DB   
	Kafka Kafka 
}
type App struct {
	Name    string 
	Version string 
}
type HTTP struct {
	Host string 
	Port string 
}
type DB struct {
	Host     string 
	Port     string 
	User     string 
	Password string 
	DBName   string 
	PgDriver string 
	Schema   string 
}
type Kafka struct {
	BootstrapServers string 
	Topic            string 
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config/config.yml"
	}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("сonfiguration error: %v", err)
	}
	return cfg, nil
}
