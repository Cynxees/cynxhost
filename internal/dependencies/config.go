package dependencies

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Key   	string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`

	Router struct {
		Default string `mapstructure:"default"`
	} `mapstructure:"router"`

	Database struct {
		MySQL struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Database string `mapstructure:"database"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Dialect  string `mapstructure:"dialect"`
			Pool     struct {
				Max     int `mapstructure:"max"`
				Min     int `mapstructure:"min"`
				Acquire int `mapstructure:"acquire"`
				Idle    int `mapstructure:"idle"`
			} `mapstructure:"pool"`
		} `mapstructure:"mysql"`

		Redis struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Password string `mapstructure:"password"`
		} `mapstructure:"redis"`

		Elasticsearch struct {
			Host string `mapstructure:"host"`
			Port int    `mapstructure:"port"`
			Log  string `mapstructure:"log"`
		} `mapstructure:"elasticsearch"`

		RabbitMQ struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"rabbitmq"`
	} `mapstructure:"database"`

	Logging struct {
		Level        string   `mapstructure:"level"`
		Format       string   `mapstructure:"format"`
		Destinations []string `mapstructure:"destinations"`
	} `mapstructure:"logging"`

	Aws struct {
		AccessKeyId        string   `mapstructure:"accessKeyId"`
		AccessKeySecret    string   `mapstructure:"accessKeySecret"`
	} `mapstructure:"aws"`

	Security struct {
		JWT struct {
			Secret    string `mapstructure:"secret"`
			ExpiresIn string `mapstructure:"expiresIn"`
		} `mapstructure:"jwt"`

		CORS struct {
			Enabled bool   `mapstructure:"enabled"`
			Origin  string `mapstructure:"origin"`
		} `mapstructure:"cors"`
	} `mapstructure:"security"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}
