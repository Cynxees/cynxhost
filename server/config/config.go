package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {

	App struct {
		Name              string `config:"APP_NAME"`
		Key               string `config:"APP_KEY"`
		Address           string `config:"APP_ADDRESS"`
		Port              string `config:"APP_PORT"`
		MicroservicePort  string `config:"APP_MICROSERVICE_PORT"`
	}

	Db struct {
		Port     string `config:"DB_PORT"`
		Host     string `config:"DB_HOST"`
		Name     string `config:"DB_NAME"`
		Username string `config:"DB_USERNAME"`
		Password string `config:"DB_PASSWORD"`
	}

	Aws struct {
		AccessKeyId     string `config:"AWS_ACCESS_KEY_ID"`
		AccessKeySecret string `config:"AWS_ACCESS_KEY_SECRET"`
	}

	Microservice struct {
		Ip struct {
			Host string `config:"MICROSERVICE_IP_HOST"`
			Port    string `config:"MICROSERVICE_IP_PORT"`
		}
	}
}

func InitConfig(path string) *Config {
	if path == "" {
		godotenv.Load(".env")
	} else {
		godotenv.Load(path)
	}

	config := &Config{}
	populateConfig(config)

	fmt.Println(config)
	return config
}

func populateConfig(config interface{}) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("config")

		// If the field is a nested struct, recurse
		if field.Kind() == reflect.Struct {
			populateConfig(field.Addr().Interface())
			continue
		}

		if tag == "" {
			panic("config tag missing for field: " + fieldType.Name)
		}

		// Fetch the environment variable based on the tag
		envValue, exists := os.LookupEnv(tag)
		if !exists {
			panic("environment variable not found: " + tag)
		}
		
		// Handle string slices for comma-separated values
		if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.String {
			field.Set(reflect.ValueOf(strings.Split(envValue, ",")))
		} else if field.Kind() == reflect.String {
			field.SetString(envValue)
		}
	
	}
}

func (c *Config) AsString() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}
