package cfg

import (
	"app/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
)

type ConfigService interface {
	LoadConfig(runMode string) error
	GetAppConfig() appCfg
	GetRepoConfig() repoCfg
	GetTransportConfig() transportCfg
}

type configService struct {
	app appCfg

	logService logger.LoggerService
}

func NewConfigService(logService logger.LoggerService) ConfigService {
	return &configService{logService: logService}
}

func (cs *configService) LoadConfig(runMode string) error {
	viper.SetConfigName("config." + runMode)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		err := fmt.Errorf("cfg.configService.LoadConfig: error on load configuration, err: %v", err)

		cs.logService.Error(err.Error())

		return err
	}

	go cs.logService.Info("cfg.configService.LoadConfig: config read successfully")

	var appConfig appCfg

	if err := viper.UnmarshalKey("App", &appConfig); err != nil {
		err := fmt.Errorf("cfg.configService.LoadConfig: error on unmarsahl app structure, err: %v", err)

		cs.logService.Error(err.Error())

		return err
	}

	go cs.logService.Info("cfg.configService.LoadConfig: successfully unmarshalling config to service")

	cs.app = appConfig

	return nil
}

func (cs *configService) GetAppConfig() appCfg {
	return cs.app
}

func (cs *configService) GetRepoConfig() repoCfg {
	return cs.app.Repo
}

func (cs *configService) GetTransportConfig() transportCfg {
	return cs.app.Transport
}
