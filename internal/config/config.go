package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	Database database
	Server   server
	Admin    admin
}

type database struct {
	Name     string
	Address  string
	Username string
	Password string
	// NetworkType string
}

type server struct {
	Debug bool
}

type admin struct {
	Username string
	Password string
}

var v *viper.Viper
var c *config

func init() {
	v = viper.New()

	// Support for config files
	cp := os.Getenv("RF_CONFIG_PATH")
	v.SetConfigName("config")
	v.AddConfigPath(cp)
	v.AddConfigPath(".")

	// Support for environment variables
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.SetEnvPrefix("RF")
	v.AutomaticEnv()

	// Map environment variables to structs
	v.BindEnv("database.name", "DATABASE_NAME")
	v.BindEnv("database.address", "DATABASE_ADDRESS")
	v.BindEnv("database.username", "DATABASE_USERNAME")
	v.BindEnv("database.password", "DATABASE_PASSWORD")

	// Configure default values
	v.SetDefault("database.name", "remotefalcon")
	// v.SetDefault("database.networkType", "tcp")

	v.SetDefault("server.debug", false)

	v.ReadInConfig()
	v.WatchConfig()
}

func GetConfig() *config {
	v.Unmarshal(&c)
	return c
}
