package main

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"os"
)

type ConfigLog struct {
	Path string
	Level string
}

type ConfigServer struct {
	Username string
	Password string
	Http ConfigServerHttp
	Sock ConfigServerSock
}

type ConfigServerHttp struct {
	Enable bool
	Host string
	Port uint32
}

type ConfigServerSock struct {
	Enable bool
	File   string
	Chmod  string
	Chown  string
}

type ConfigInclude struct {
	Directory string
}

type ConfigProgram struct {
	Name string
	Group string
	Cmd string
	Args []string
	Environment map[string]string
	Autostart bool
	Instances uint
	Priority uint
	Settle uint
	Retries uint
	Delay uint
	Stopsignal string
	Stoptimeout string
	Runuser string
	Rungroup string
}

type Config struct {
	Log ConfigLog
	Server ConfigServer
	Include ConfigInclude
	Programs []ConfigProgram
}

func ReadConfig(name string, paths []string) (*Config, error) {
	var cfg Config

	vpr := viper.New()

	vpr.SetConfigName(name)

	if len(paths) < 1 {
		paths = []string{".", string(os.PathSeparator) + "etc" + string(os.PathSeparator) + "gogo"}
	}
	for _, a := range paths {
		vpr.AddConfigPath(a)
	}

	err := vpr.ReadInConfig()

	if err != nil {
		return nil, err
	}

	vpr.Unmarshal(&cfg)

	cfgDir := filepath.Dir(vpr.ConfigFileUsed())

	dir := cfg.Include.Directory
	programsDir := filepath.Clean(cfgDir + string(os.PathSeparator) + dir) + string(os.PathSeparator)
	dirList, err := ioutil.ReadDir(programsDir)
	if err != nil {
		panic(err)
	}
	for _, f := range dirList {
		if f.IsDir() {
			continue
		}
		pCfg := ConfigProgram{}
		pPath := programsDir + f.Name()
		pVpr := viper.New()
		pVpr.SetConfigFile(pPath)
		pVpr.ReadInConfig()
		pVpr.Unmarshal(&pCfg)
		cfg.Programs = append(cfg.Programs, pCfg)
	}

	return &cfg, nil
}
