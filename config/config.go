package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Baidu struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type Tencent struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func TencentConfig() (*Tencent, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	tencent := &Tencent{
		Key:    viper.Get("tencent.key").(string),
		Secret: viper.Get("tencent.secret").(string),
	}
	return tencent, nil
}
