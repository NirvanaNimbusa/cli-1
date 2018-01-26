package config

import (
	"flag"
	"os"
	"path/filepath"

	C "github.com/ActiveState/ActiveState-CLI/internal/constants"
	"github.com/ActiveState/ActiveState-CLI/internal/print"
	"github.com/shibukawa/configdir"
	"github.com/spf13/viper"
)

var configNamespace = C.ConfigNamespace
var configDirs configdir.ConfigDir
var configDir *configdir.Config

var exit = os.Exit

func init() {
	if flag.Lookup("test.v") == nil {
		Init()
	}
}

func Init() {
	// Prepare our config dir, eg. ~/.config/activestate/cli
	configDirs = configdir.New(configNamespace, "cli")
	configDirs.LocalPath, _ = filepath.Abs(".")
	configDir = configDirs.QueryFolders(configdir.Global)[0]

	if !configDir.Exists(C.ConfigFileName) {
		configDir.Create(C.ConfigFileName)
	}

	// Prepare viper, which is a library that automates configuration
	// management between files, env vars and the CLI
	viper.SetConfigName(C.ConfigName)
	viper.SetConfigType(C.ConfigFileType)
	viper.AddConfigPath(configDir.Path)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		print.Error("Can't read config: %s", err)
		exit(1)
	}
}

// Save the config state to the config file
func Save() {
	viper.WriteConfig()
}