package cfg

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigService interface {
	LoadConfig(runMode string) error
	GetAppConfig() appCfg
	GetDBConfig() repoCfg
}

type configService struct {
	app appCfg
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (cs *configService) LoadConfig(runMode string) error {
	viper.SetConfigName("config." + runMode)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cfg: error on load configuration, err: %v", err)
	}

	var appConfig appCfg

	if err := viper.UnmarshalKey("App", &appConfig); err != nil {
		return fmt.Errorf("cfg: error on unmarsahl app structure, err: %v", err)
	}

	cs.app = appConfig

	return nil
}

func (cs *configService) GetAppConfig() appCfg {
	return cs.app
}

func (cs *configService) GetDBConfig() repoCfg {
	return cs.app.Repo
}
