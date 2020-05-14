package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Global *Config
)

//Config 設定組態
type Config struct {
	Title   string  `toml:"title"`
	Discord Discord `toml:"discord"`
}

//Discord Discord
type Discord struct {
	WebhookLink string `toml:"webhookLink"`
}

func (a Discord) info() string {
	return a.WebhookLink
}

func init() {
	println(" < - Config loading - > ")
	LoadGlobalConfig("config/config.toml")
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
