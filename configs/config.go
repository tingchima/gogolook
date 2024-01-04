// Package configs provides
package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var base string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	base = filepath.Dir(currentFile)
}

type AppConfig struct {
	Viper *viper.Viper
}

func NewConfig(name string) *AppConfig {

	configPaths := []string{base}

	v := viper.New()

	v.SetConfigName(name)   // name of config file (without extension)
	v.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	for i := range configPaths {
		v.AddConfigPath(configPaths[i])
	}
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		log.Printf("config file: %s", err)
	}

	return &AppConfig{Viper: v}
}

// Database .
type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

func (c *AppConfig) Database() *Database {
	return &Database{
		Host:     c.Viper.GetString("database.host"),
		Port:     c.Viper.GetString("database.port"),
		Username: c.Viper.GetString("database.username"),
		Password: c.Viper.GetString("database.password"),
		DBName:   c.Viper.GetString("database.db_name"),
	}
}
