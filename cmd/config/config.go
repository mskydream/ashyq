package config

import "github.com/spf13/viper"

type Configuration struct {
	Database DB     `mapstructure:"db"`
	Port     string `mapstructure:"port"`
	Url      string `mapstructure:"url"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Name     string `mapstructure:"name"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigName("prod")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
