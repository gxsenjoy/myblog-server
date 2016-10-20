package myblog

import "github.com/spf13/viper"

type Config map[string]interface{}

func ParseConfig(defaultConfig Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.myblog")
	viper.AddConfigPath(".")

	for key, defaultValue := range defaultConfig {
		viper.Set(key, defaultValue)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
