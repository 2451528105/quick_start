package ntool

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

func LoadConfigFromFile(filePath string, v any, listen bool) {
	vp := viper.New()
	if err := load(vp, filePath, v); err != nil {
		fmt.Println("load config file failed, err:", err)
		return
	}
	if listen {
		vp.WatchConfig()
		vp.OnConfigChange(func(_ fsnotify.Event) {
			if err := load(vp, filePath, v); err != nil {
				fmt.Println("reload config file failed, err:", err)
			} else {
				fmt.Println("reload config file success")
			}
		})
	}
}

func load(vp *viper.Viper, filePath string, v any) error {
	vp.SetConfigFile(filePath)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	// return vp.Unmarshal(v)
	return vp.Unmarshal(v, func(dc *mapstructure.DecoderConfig) { dc.TagName = "yaml" })

}
