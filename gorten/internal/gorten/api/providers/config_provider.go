package providers

import (
	"gorten/pkg/logs"

	"gorten/internal/gorten/config"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func ConfigProvider() *config.AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logs.Logger.Fatalf("Failed to read config file: %v", err)
	}

	var appConfig config.AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		logs.Logger.Fatalf("Failed to unmarshal config: %v", err)
	}

	return &appConfig
}

var ConfigModule = fx.Options(
	fx.Provide(ConfigProvider),
)
