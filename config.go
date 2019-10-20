package main

import (
	"log"
	"os"
    "path/filepath"

	"github.com/BurntSushi/toml"
)

const FileConfig = "/config/config.toml"

var _config *Config = nil

type Config struct {
	LDAP LDAPConfig
}
type LDAPConfig struct {
	ServerHost     string
	ServerPort     uint16
	IsTLS          bool
	CACertFilePath string
	BindDN         string
	BindPassword   string
	BaseDN         string
	SearchFilter   string
}

func LoadConfig() *Config {
	if _config == nil {
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
        exePath := filepath.Dir(exe)

		if _, err := toml.DecodeFile(exePath+FileConfig, &_config); err != nil {
			log.Fatal(err)
		}
	}
	return _config
}
