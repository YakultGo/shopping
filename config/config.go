package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Mysql string `mapstructure:"mysql"`
		Redis string `mapstructure:"redis"`
	} `mapstructure:"db"`
	User struct {
		Http struct {
			Host        string `mapstructure:"host"`
			Port        int    `mapstructure:"port"`
			ServiceName string `mapstructure:"service_name"`
		} `mapstructure:"http"`
		Grpc struct {
			Host        string   `mapstructure:"host"`
			ServiceName string   `mapstructure:"service_name"` // 带下划线的字段要是用mapstructure标签
			Tags        []string `mapstructure:"tags"`
		} `mapstructure:"grpc"`
	} `mapstructure:"user"`
	Sms struct {
		Grpc struct {
			Host        string   `mapstructure:"host"`
			ServiceName string   `mapstructure:"service_name"` // 带下划线的字段要是用mapstructure标签
			Tags        []string `mapstructure:"tags"`
		} `mapstructure:"grpc"`
	} `mapstructure:"sms"`
	Consul struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"consul"`
}

var Conf = new(Config)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
	return nil
}
