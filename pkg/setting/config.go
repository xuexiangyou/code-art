package setting

import (
	"github.com/spf13/viper"
	"github.com/xuexiangyou/code-art/config"
	"os"
)

const (
	EnvValDev  = "dev"
	EnvValTest = "test"
	EnvValProd = "prod"

	ConfigFilePathDev  = "./config/dev/config.yaml"
	ConfigFilePathTest = "./config/test/config.yaml"
	ConfigFilePathProd = "./config/prod/config.yaml"
)

//NewConfig 初始化配置
func NewConfig() (*config.Config, error) {
	configFile := getConfigFilePah()

	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

//getConfigFilePah 更新环境变量获取配置文件路径
func getConfigFilePah() string {
	var configPath string
	envVal := os.Getenv("ENV_CONF")

	switch envVal {
	case EnvValDev:
		configPath = ConfigFilePathDev
	case EnvValTest:
		configPath = ConfigFilePathTest
	case EnvValProd:
		configPath = ConfigFilePathProd
	}
	return configPath
}
