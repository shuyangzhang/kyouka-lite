package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	Token string `mapstructure:"TOKEN"`

	Prefix        string   `mapstructure:"PREFIX"`
	PrefixAliases []string `mapstructure:"PREFIX_ALIASES"`
	AllPrefixes   []string
}

func initEnvConfigs() {
	EnvConfigs = loadEnvVariables()
	EnvConfigs.AllPrefixes = append(EnvConfigs.PrefixAliases, EnvConfigs.Prefix)
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed to load .env file, err: %v \n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Failed to unmarshal .env file, err: %v \n", err)
	}

	return
}
