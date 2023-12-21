package load

import (
	"github.com/Didjacome/varacode-api/schemas"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *schemas.Veracode

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	keys := []string{"Credetial.apiKeyID", "Credetial.apiKeySecret", "Credetial.apiUrl"}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			for _, key := range keys {
				if !viper.IsSet(key) {
					log.Fatalf("Error: %v \n Environment Not Foud: %v", err, key)
				}
			}
		} else {
			log.Fatalf("Error: %v", err)
		}
	}

	config = new(schemas.Veracode)

	config.ApiCredetial = schemas.Auth{
		ApiKeyID:     viper.GetString("Credetial.apiKeyID"),
		ApiKeySecret: viper.GetString("Credetial.apiKeySecret"),
		ApiUrl:       viper.GetString("Credetial.apiUrl"),
	}

	return nil
}

func GetApiCredetial() schemas.Auth {
	return config.ApiCredetial
}