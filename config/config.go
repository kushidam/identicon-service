package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	NAMESPACE   string
	BLOCKS_SIZE int
	DENSITY     int
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	Config = ConfigList{
		NAMESPACE:   cfg.Section("identicon").Key("namespace").String(),
		BLOCKS_SIZE: cfg.Section("identicon").Key("blocks_size").MustInt(),
		DENSITY:     cfg.Section("identicon").Key("density").MustInt(),
	}
}
