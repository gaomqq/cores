package config

import (
	"github.com/spf13/viper"
)

type config struct {
	DataID string
	Group  string
	Ip     string
	Port   int
}

func Viper() config {

	v := viper.New()
	v.SetConfigFile("/Users/gaomengwei/go/src/project/core/config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var c config
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}

	return c

}
