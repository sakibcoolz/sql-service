package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Credentials struct {
	Database string
	Host     string
	Port     string
	User     string
	Password string
}

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	name    string
	Log     *zap.Logger
	Creden  Credentials
	Service ServiceConfig
}

func NewConfigSetup(Log *zap.Logger, ServiceName string) *Config {
	return &Config{
		name: ServiceName,
		Log:  Log,
	}
}

func (c *Config) EnvConfiguration() {
	if c.Creden.Database = os.Getenv("DATABASE"); c.Creden.Database == "" {
		c.Log.Fatal("Database environment variable not set")
	}

	if c.Creden.Host = os.Getenv("HOST"); c.Creden.Host == "" {
		c.Log.Fatal("HOST environment variable not set")
	}

	if c.Creden.Port = os.Getenv("PORT"); c.Creden.Port == "" {
		c.Log.Fatal("PORT environment variable not set")
	}

	if c.Creden.User = os.Getenv("USER"); c.Creden.User == "" {
		c.Log.Fatal("USER environment variable not set")
	}

	if c.Creden.Password = os.Getenv("PASSWORD"); c.Creden.Password == "" {
		c.Log.Fatal("PASSWORD environment variable not set")
	}

	if c.Service.Host = os.Getenv("SERVICEHOST"); c.Service.Host == "" {
		c.Log.Fatal("SERVICEHOST environment variable not set")
	}

	if c.Service.Port, _ = strconv.Atoi(os.Getenv("SERVICEPORT")); c.Service.Port == 0 {
		c.Log.Fatal("SERVICEPORT environment variable not set")
	}
}

func (c *Config) GetConfiguration() {
	c.EnvConfiguration()
}
