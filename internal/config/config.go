package config

var NCfg *NacosConfig = &NacosConfig{}
var Cfg *Config = &Config{}

type NacosConfig struct {
	Nacos struct {
		IpAddr         string `yaml:"IpAddr"`
		Port           uint64 `yaml:"Port"`
		Path           string `yaml:"Path"`
		DataId         string `yaml:"DataId"`
		Group          string `yaml:"Group"`
		Namespace      string `yaml:"Namespace"`
		RequestTimeout uint64 `yaml:"RequestTimeout"`
		LogDir         string `yaml:"LogDir"`
		CacheDir       string `yaml:"CacheDir"`
		LogLevel       string `yaml:"LogLevel"`
	} `yaml:"Nacos"`
}

type NacosConfigInfo struct {
	IpAddr         string `yaml:"IpAddr"`
	Port           string `yaml:"Port"`
	Path           string `yaml:"Path"`
	DataId         string `yaml:"DataId"`
	Group          string `yaml:"Group"`
	Namespace      string `yaml:"Namespace"`
	RequestTimeout string `yaml:"RequestTimeout"`
	LogDir         string `yaml:"LogDir"`
	CacheDir       string `yaml:"CacheDir"`
	LogLevel       string `yaml:"LogLevel"`
}

type Config struct {
	GameId     int               `yaml:"GameId"`
	ServerName string            `yaml:"ServerName"`
	Env        string            `yaml:"Env"`
	WsPath     string            `yaml:"WsPath"`
	Port       string            `yaml:"Port"`
	Log        *LogConfig        `yaml:"Log"`
	Redis      *RedisConfig      `yaml:"Redis"`
	RocketMq   *RocketMqConfig   `yaml:"RocketMq"`
	Rpc        *RpcConfig        `yaml:"Rpc"`
	SkyWalking *SkyWalkingConfig `yaml:"SkyWalking"`
}

type LogConfig struct {
	Level string      `yaml:"Level"`
	Mode  string      `yaml:"Mode"`
	File  *FileConfig `yaml:"File"`
}

type FileConfig struct {
	FilePath   string `yaml:"FilePath"`
	MaxSize    int    `yaml:"MaxSize"`
	MaxBackups int    `yaml:"MaxBackup"`
	MaxAge     int    `yaml:"MaxAge"`
	Compress   bool   `yaml:"Compress"`
	LocalTime  bool   `yaml:"LocalTime"`
}

type RedisConfig struct {
	Addr       string `yaml:"Addr"`
	ClientName string `yaml:"ClientName"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	DB         int    `yaml:"DB"`
}

type RocketMqConfig struct {
	Addr      string `yaml:"Addr"`
	NameSpace string `yaml:"NameSpace"`
	Group     string `yaml:"Group"`
	AccessKey string `yaml:"AccessKey"`
	SecretKey string `yaml:"SecretKey"`
}

type RpcConfig struct {
	Name      string `yaml:"Name"`
	Addr      string `yaml:"Addr"`
	NameSpace string `yaml:"NameSpace"`
	Group     string `yaml:"Group"`
	Filters   string `yaml:"Filters"`
}

type SkyWalkingConfig struct {
	Enable      bool     `yaml:"Enable"`      // 是否启用链路追踪
	ServiceName string   `yaml:"ServiceName"` // 服务名称
	Addr        string   `yaml:"Addr"`        // OAP 地址
	BlackList   []string `yaml:"BlackList"`   // 黑名单
}
