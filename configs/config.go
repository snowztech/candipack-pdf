package configs

import "github.com/spf13/viper"

type Config struct {
	Port    int
	APIKey	string
}

func Load() Config {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	conf := Config{}
	conf.Port = getIntWithDefault("PORT", 9000)
	conf.APIKey = viper.GetString("API_KEY")

	return conf
}

func getIntWithDefault(key string, defaultValue int) int {
	value := viper.GetInt(key)
	if value == 0 {
		return defaultValue
	}
	return value
}