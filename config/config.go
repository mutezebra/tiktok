package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	System  *System             `yaml:"system"`
	Service map[string]*Service `yaml:"service"`
	Mysql   *Mysql              `yaml:"mysql"`
	Etcd    *Etcd               `yaml:"etcd"`
}

type System struct {
	OS           string `yaml:"os"`
	Status       string `yaml:"status"`
	WorkerID     int64  `yaml:"worker_id" mapstructure:"worker_id"`
	DataCenterID int64  `yaml:"data_center_id" mapstructure:"data_center_id"`
}

type Service struct {
	ServiceName string `yaml:"service_name" mapstructure:"service_name"`
	Address     string `yaml:"address"`
}

type Mysql struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}

type Etcd struct {
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	Endpoint      string `yaml:"endpoint"`
	ServicePrefix string `yaml:"service_prefix" mapstructure:"service_prefix"`
}

// InitConfig initializes the configuration for the project
// and unmarshals the configuration into the global variable "Conf"
func InitConfig() {
	wd, _ := os.Getwd()
	viper.AddConfigPath(wd + "/config/")   // linux
	viper.AddConfigPath(wd + "\\config\\") // linux
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}
