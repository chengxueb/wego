package wego

import (
	"github.com/wego/config"
	"github.com/wego/utils"
	"os"
	"path/filepath"
	"time"
)

var (
	WConfig           *WegoConfig
	AppPath           string
	AppConfigFilename string
)

func init() {
	WConfig = NewConfig()
	var err error
	AppPath, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	AppConfigFilename = filepath.Join(AppPath, WConfig.AppConfig.ConfigDir, WConfig.AppConfig.ConfigFilename)
	if !utils.FileExists(AppConfigFilename) {
		return
	}

	WConfig.UserAppConfig = config.ParseConfigFile(AppConfigFilename)
}

func NewConfig() *WegoConfig {
	a := NewAppConfig()
	w := NewWebConfig()
	return &WegoConfig{WebConfig: w, AppConfig: a}
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AppName:        "wego",
		ConfigFilename: "app.conf",
		ConfigDir:      "config",
		RoutersDir:     "routers",
		ControllersDir: "controllers",
		ViewsDir:       "views",
	}
}

func NewWebConfig() *WegoWebConfig {
	return &WegoWebConfig{
		Addr:         ":8080",
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
}

type WegoConfig struct {
	AppConfig     *AppConfig
	WebConfig     *WegoWebConfig
	UserAppConfig map[string]string
}

type AppConfig struct {
	AppName         string
	ConfigDir       string
	ConfigFilename  string
	RoutersDir      string
	RoutersFilename string
	ControllersDir  string
	ViewsDir        string
}

type WegoWebConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
