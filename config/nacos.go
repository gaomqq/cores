package config

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type T struct {
	App struct {
		Ip     string `json:"ip"`
		Port   string `json:"port"`
		Secret string `json:"secret"`
	} `json:"app"`
	Mysql struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`
	Consul struct {
		Ip   string `json:"ip"`
		Port string `json:"port"`
	} `json:"consul"`
}

func ServiceNaCos() (T, error) {

	v := Viper()

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: v.Ip,
			Port:   uint64(v.Port),
		},
	}

	// 创建动态配置客户端
	namingClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, err := namingClient.GetConfig(vo.ConfigParam{
		DataId: v.DataID,
		Group:  v.Group})
	if err != nil {
		return T{}, err
	}

	var t T

	json.Unmarshal([]byte(content), &t)
	return t, nil

}
