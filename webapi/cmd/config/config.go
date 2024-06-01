package config

import (
	_ "webapi/pkg/config"

	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

func init() {
	// App
	viper.SetDefault("API_PORT", "")
	viper.SetDefault("APP_ENV", "production")
	viper.SetDefault("API_JWT_SECRET", "")
	// DB
	viper.SetDefault("DB_DNS", "")
	viper.SetDefault("DB_MAX_IDLE_CONNS", 10)
	viper.SetDefault("DB_MAX_OPEN_CONNS", 10)

	viper.AutomaticEnv()
	pretty.Log("Viper initialized:\n\n", viper.AllSettings())
}
