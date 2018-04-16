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

type ConfigTasks struct {
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
	Tasks []ConfigTasks
}

const DefaultLogPath = "gogo.log"
const DefaultLogLevel = "info"
const DefaultServerUsername = "gogo"
const DefaultServerPassword = "{sha256}65B63136D9BA9D576B7DEF9DE5B767A1B4F513A992F316A50B3F548375638078"
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
				var subCfg Config
				_, err := toml.DecodeFile(f, &subCfg)
				if err != nil {
					panic(err)
				}
				for _, v := range subCfg.Tasks {
					cfg.Tasks = append(cfg.Tasks, v)
				}
			}
		}
	}
	return cfg, nil
}
