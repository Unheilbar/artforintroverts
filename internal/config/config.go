package config

import "github.com/spf13/viper"

type Config struct {
	DSN                  string `mapstructure:"DSN"`
	DbName               string `mapstructure:"DATABASE_NAME"`
	CollectionName       string `mapstructure:"COLLECTION_NAME"`
	ServerAddress        string `mapstructure:"SERVER_ADDRESS"`
	CacheRefreshInterval uint   `mapstructure:"CACHE_REFRESH_INTERVAL_MS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
