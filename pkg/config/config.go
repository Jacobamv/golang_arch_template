package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// ImportConfigs find config file and creates a new Config struct
func ImportConfigs() (config *Config) {
	godotenv.Load("../.env")

	result := GetKeysOfStruct(Config{}, "")
	viper := viper.New()

	for _, v := range result {
		val := os.Getenv(v)

		key := strings.ReplaceAll(v, "_", ".")
		viper.Set(key, val)
	}

	viper.Unmarshal(&config)

	return
}

func GetKeysOfStruct(obj any, parent_name string) []string {
	v := reflect.ValueOf(obj)
	typeOfS := v.Type()
	keys := make([]string, 0)
	var new_keys []string
	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Type.Kind() != reflect.Struct {
			keys = append(keys, parent_name+typeOfS.Field(i).Tag.Get("mapstructure"))
		} else {

			if parent_name == "" {
				new_keys = GetKeysOfStruct(v.Field(i).Interface(), typeOfS.Field(i).Tag.Get("mapstructure")+"_")
			} else {
				new_keys = GetKeysOfStruct(v.Field(i).Interface(), parent_name+typeOfS.Field(i).Tag.Get("mapstructure")+"_")
			}

			for _, j := range new_keys {
				keys = append(keys, j)
			}
		}
	}

	return keys
}
