package cargo

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func ReadInPackageName(cratePath string) (string, error) {
	v := viper.New()
	v.SetConfigFile(filepath.Join(cratePath, "Cargo.toml"))
	if err := v.ReadInConfig(); err != nil {
		return "", fmt.Errorf("fail to read Cargo's manifest: %w", err)
	}

	if name := v.GetString("lib.name"); name != "" {
		return name, nil
	}

	name := v.GetString("package.name")
	if name == "" {
		return "", errors.New("missing package name")
	}

	return name, nil
}

func UseLocalCrate4Dependency(cratePath, dependency, localCratePath string) error {
	localCratePkgName, err := ReadInPackageName(localCratePath)
	if err != nil {
		return fmt.Errorf("fail read package name of local crate: %w", err)
	}

	localCrateAbsPath, err := filepath.Abs(localCratePath)
	if err != nil {
		return fmt.Errorf("fail to get the absolute path of the local crate: %w", err)
	}

	v := viper.New()
	v.SetConfigFile(filepath.Join(cratePath, "Cargo.toml"))
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// @TODO: validate format, such if dependency exists
	pkgKey := fmt.Sprintf("dependencies.%s.package", dependency)
	v.Set(pkgKey, localCratePkgName)

	pathKey := fmt.Sprintf("dependencies.%s.path", dependency)
	v.Set(pathKey, localCrateAbsPath)

	if err := v.WriteConfig(); err != nil {
		return fmt.Errorf("fail to renew config: %w", err)
	}

	return nil
}
