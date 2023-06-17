package utility

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	MigrationPath     string `mapstructure:"MIGRATION_PATH"`
}

type Environment string

const (
	DEV  Environment = "dev"
	TEST             = "test"
	PROD             = "prod"
)

func NewConfig(environment Environment) (config Config, err error) {
	viper.AddConfigPath(".")
	switch environment {
	case DEV:
		viper.SetConfigName("app_dev")
	case TEST:
		viper.SetConfigName("app_test")
	case PROD:
		viper.SetConfigName("app_prod")
	}
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error loading config")
		fmt.Println(err)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
