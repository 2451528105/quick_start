package config

import (
	"fmt"
	"quick-start/pkg/ntool"

	"github.com/ivy-mobile/odin/encoding/yaml"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func Init() {
	readNacosConfig()
	readConfigFromNacos()

}

// readNacosConfig 读取本地nacos配置
func readNacosConfig() {
	ntool.LoadConfigFromFile("config/config.yaml", NCfg, true)
	fmt.Println(NCfg.Nacos)
}

// readConfigFromNacos 读取nacos中的游戏配置
func readConfigFromNacos() {
	clientConfig := constant.NewClientConfig(
		constant.WithNamespaceId(NCfg.Nacos.Namespace),
		constant.WithTimeoutMs(NCfg.Nacos.RequestTimeout),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(NCfg.Nacos.LogDir),
		constant.WithCacheDir(NCfg.Nacos.CacheDir),
		constant.WithLogLevel(NCfg.Nacos.LogLevel),
	)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      NCfg.Nacos.IpAddr,
			Port:        NCfg.Nacos.Port,
			ContextPath: NCfg.Nacos.Path,
		},
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		fmt.Println("new config client failed, err:", err)
		return
	}
	// 获取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: NCfg.Nacos.DataId,
		Group:  NCfg.Nacos.Group,
	})
	if err != nil {
		fmt.Println("get config failed, err:", err)
		return
	}
	// 解析配置
	if err := yaml.Unmarshal([]byte(content), Cfg); err != nil {
		fmt.Println("unmarshal config failed, err:", err)
		return
	}

}
