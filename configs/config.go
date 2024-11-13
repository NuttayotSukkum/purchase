package configs

import (
	"context"
	"github.com/NuttayotSukkum/purchase/internal/constants"
	logger "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	App     App     `mapstructure:"app"`
	Secrets Secrets `mapstructure:"secrets"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
	env     string `mapstructure:"env"`
}

type Secrets struct {
	CloudSqlHost   string `mapstructure:"cloud-sql-gormhost"`
	CloudSqlPort   string `mapstructure:"cloud-sql-port"`
	CloudSqlUser   string `mapstructure:"cloud-sql-username"`
	CloudSqlPass   string `mapstructure:"cloud-sql-password"`
	CloudSqlDBName string `mapstructure:"cloud-sql-dbname"`
}

var config *Config
var configOne sync.Once

func InitConfig(ctx context.Context) *Config {
	configOne.Do(func() {
		viper.AddConfigPath(constants.ConfigPath)
		viper.SetConfigName(constants.ConfigName)
		viper.SetConfigType(constants.ConfigType)
		//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			logger.Infof("%s config file not found. using default/env config: %s", ctx, err)
		}
		config = &Config{}
		if err := viper.Unmarshal(&config); err != nil {
			logger.Errorf("%s unable to decode struct: %s", ctx, err)
			return
		} else {
			logger.Infof("%s config: %+v", ctx, config)
		}
	})
	return config
}
