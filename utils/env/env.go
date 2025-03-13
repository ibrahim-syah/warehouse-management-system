package envutils

import (
	"warehouse-management-system/utils/loggerutils"

	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		loggerutils.LoggerSingleton.Fatal("Can't find the file .env : ", err)
	}
}
