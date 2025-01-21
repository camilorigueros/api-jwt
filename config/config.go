package config

import (
	"github.com/spf13/viper"
	"log"
)

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int64  `mapstructure:"expiration"`
	Issuer     string `mapstructure:"issuer"`
}

type Config struct {
	Security struct {
		JWT   JWTConfig         `mapstructure:"jwt"`
		Users map[string]string `mapstructure:"users"`
	} `mapstructure:"security"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")   // Nombre del archivo sin extensión
	viper.SetConfigType("yaml")     // Tipo de archivo
	viper.AddConfigPath("./config") // Directorio actual

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer la configuración: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error al parsear la configuración: %v", err)
	}
}
