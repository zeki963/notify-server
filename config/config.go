package config

import (
	"github.com/BurntSushi/toml"
	"github.com/zorhayashi/notify-server/util"
)

var (
	//Global 全域配置
	Global *Config
)

//Config toml
type Config struct {
	Title   string `toml:"Title"`
	Systeam struct {
		Admin bool `toml:"Admin"`
	} `toml:"systeam"`
	Discord struct {
		DiscordStatus bool   `toml:"DiscordStatus"`
		WebhookLink   string `toml:"WebhookLink"`
	} `toml:"Discord"`
	Line struct {
		LineStatus         bool   `toml:"LineStatus"`
		ChannelSecret      string `toml:"channelSecret"`
		ChannelAccessToken string `toml:"channelAccessToken"`
	} `toml:"Line"`
}

//Configinit 初始化設定參數
func Configinit() {
	util.Info("Config loding..")
	if err := LoadGlobalConfig("config/config.toml"); err != nil {
		util.Error("Config Errorrrrr")
	}
	util.Success(`Config.Title [` + Global.Title + `]`)
}

//LoadGlobalConfig 加載全局配置
func LoadGlobalConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	Global = c
	return nil
}

//GetGlobalConfig 獲取全局配置
func GetGlobalConfig() *Config {
	if Global == nil {
		return &Config{}
	}
	return Global
}

//ParseConfig 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
