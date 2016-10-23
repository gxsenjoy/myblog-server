package myblog

import "github.com/spf13/viper"

type Config map[string]interface{}

func InitConfig(override Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.myblog")
	viper.AddConfigPath(".")

	for key, val := range override {
		if val.(string) != "" {
			viper.Set(key, val)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
