package load

import (
	"strings"
	"github.com/Didjacome/varacode-api/schemas"
	"github.com/spf13/viper"
)
var config *schemas.Veracode

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	config = new(schemas.Veracode)

	config.ApiCredetial = schemas.Auth{
		ApiKeyID:     viper.GetString("Credetial.apiKeyID"),
		ApiKeySecret: viper.GetString("Credetial.apiKeySecret"),
		ApiUrl:       viper.GetString("Credetial.apiUrl"),
	}

	return err
}


func GetApiCredetial() schemas.Auth {
	return config.ApiCredetial
}