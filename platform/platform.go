package platform

import (
	"os"
	"path/filepath"
	"strings"
)

type EnvFlag struct {
	Name    string
	AltName string
}

func NormalizeEnvName(name string) string {
	return strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(name)), ".", "_")
}

func (f EnvFlag) GetValue(defaultValue func() string) string {
	if v, found := os.LookupEnv(f.Name); found {
		return v
	}
	if len(f.Name) > 0 {
		if v, found := os.LookupEnv(f.AltName); found {
			return v
		}
	}
	return defaultValue()

}

func getExecutableDir() string {
	exec, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exec)
}

func NewEnvFlag(name string) EnvFlag {
	return EnvFlag{
		Name:    name,
		AltName: NormalizeEnvName(name),
	}

}

func GetConfDirPath() string {
	const name = "v2ray.location.confdir"
	configPath := NewEnvFlag(name).GetValue(func() string { return "" })
	return configPath

}

func GetConfigurationPath() string {
	const name = "v2ray.location.config"
	configPath := NewEnvFlag(name).GetValue(getExecutableDir)
	return filepath.Join(configPath, "config.json")

}
