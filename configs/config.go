package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

const (
	ConfigName = "app_config"
	ConfigType = "env"
	ConfigFile = ".env"
)

var config *Config

type DatabaseConfig struct {
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
}

type ServerHttpConfig struct {
	Port         string `mapstructure:"PORT"`
	JwtSecret    string `mapstruct:"JWT_SECRET"`
	JwtExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

type Config struct {
	Database DatabaseConfig   `mapstructure:",squash"`
	Server   ServerHttpConfig `mapstructure:",squash"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.SetConfigFile(ConfigFile)
	viper.AutomaticEnv()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	config.Server.TokenAuth = jwtauth.New("HS256", []byte(config.Server.JwtSecret), nil)

	return config, nil
}
