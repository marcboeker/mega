package config

import (
	"flag"

	"github.com/spf13/viper"
)

// Config contains a reference to the viper configuration management.
type Config struct {
	*viper.Viper
}

// New tries to resolve and parse the YAML config file either in the current
// working dir or via command line argument.
func New() (*Config, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path of configuration file")
	flag.Parse()

	v := viper.New()
	if len(configPath) > 0 {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("../../")
	}
	setDefault(v)
	return &Config{Viper: v}, v.ReadInConfig()
}

func setDefault(v *viper.Viper) {
	// DB
	v.SetDefault("db.debugging", true)
	v.SetDefault("db.driverName", "sqlite3")
	v.SetDefault("db.dataSourceName", ":memory:?_foreign_keys=on")
	// Server
	v.SetDefault("server.logging", true)
	v.SetDefault("server.address", "localhost:8080")
	v.SetDefault("server.recoveryMiddleware", true)
	v.SetDefault("server.enablePlayground", true)
	v.SetDefault("server.playgroundPath", "/playground")
	v.SetDefault("server.queryPath", "/query")
}

// Get returns the current instance of the config.
func Get() *Config {
	return &Config{Viper: viper.GetViper()}
}
