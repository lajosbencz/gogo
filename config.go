package main

import (
	"github.com/BurntSushi/toml"
	"path"
	"os"
	"io/ioutil"
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
	Host string
	Port uint32
}

type ConfigServerSock struct {
	File   string
	Chmod  string
	Chown  string
}

type ConfigTaskShared struct {
	Environment map[string]string
	Autostart bool
	Instances uint
	Priority uint
	Settle uint
	Retries uint
	Delay uint
	Stopsignal string
	Stoptimeout uint
	Runuser string
	Rungroup string
}

type ConfigTaskItem struct {
	ConfigTaskShared
	Info string
	Group string
	Cmd string
	Args []string
}

type ConfigTask struct {
	ConfigTaskShared
	Subdir string
}

type Config struct {
	Log ConfigLog
	Server ConfigServer
	Task ConfigTask
	Tasks map[string]ConfigTaskItem
}

const DefaultLogPath = "gogo.log"
const DefaultLogLevel = "info"
const DefaultServerUsername = "gogo"
const DefaultServerPassword = "{sha256}2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b"
const DefaultServerHttpHost = "localhost"
const DefaultServerHttpPort = 8181

func defaultConfig(c *Config) {
	c.Log = ConfigLog{
		Path: DefaultLogPath,
		Level: DefaultLogLevel,
	}
	c.Server = ConfigServer{
		Username: DefaultServerUsername,
		Password: DefaultServerPassword,
		Http: ConfigServerHttp{
			Host: DefaultServerHttpHost,
			Port: DefaultServerHttpPort,
		},
	}
	c.Tasks = make(map[string]ConfigTaskItem)
}

func ReadConfig(filePath string) (Config, error) {
	cfg := Config{}
	defaultConfig(&cfg)
	_, err := toml.DecodeFile(filePath, &cfg)
	if err != nil {
		panic(err)
	}
	cfgPath := path.Dir(filePath)
	if cfg.Task.Subdir != "" {
		d := path.Clean(cfgPath + string(os.PathSeparator) + cfg.Task.Subdir)
		if DirExists(d) {
			dirList, err := ioutil.ReadDir(d)
			if err != nil {
				panic(err)
			}
			for _, l := range dirList {
				if l.IsDir() {
					continue
				}
				f := d + string(os.PathSeparator) + l.Name()
				if !FileExists(f) {
					continue
				}
				if path.Ext(f) != ".toml" {
					continue
				}
				var subCfg Config
				_, err := toml.DecodeFile(f, &subCfg)
				if err != nil {
					panic(err)
				}
				for k, v := range subCfg.Tasks {
					cfg.Tasks[k] = v
				}
			}
		}
	}
	return cfg, nil
}
