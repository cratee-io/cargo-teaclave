package config

import (
	"os/user"
	"path/filepath"
)

var homeDir string

//func DefaultConfigPath() string {
//	return filepath.Join(homeDir, "config.toml")
//}

func DefaultRegistryBaseDir() string {
	return filepath.Join(homeDir, "registry")
}

func init() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir = filepath.Join(user.HomeDir, ".cargo-teaclave")
}
