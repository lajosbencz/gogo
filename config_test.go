package main

import (
	"testing"
	"os"
	"strconv"
)

func TestDefaultConfig(t *testing.T) {
	var cfg Config
	var err error
	if cfg, err = ReadConfig("./examples/config/empty.toml"); err != nil {
		t.Fatal(err.Error())
	}
	if cfg.Log.Path != DefaultLogPath {
		t.Fatal("log.path")
	}
	if cfg.Log.Level != DefaultLogLevel {
		t.Fatal("log.level")
	}
	if cfg.Server.Username != DefaultServerUsername {
		t.Fatal("log.server.username")
	}
	if cfg.Server.Password != DefaultServerPassword {
		t.Fatal("log.server.password")
	}
	if cfg.Server.Http.Host != DefaultServerHttpHost {
		t.Fatal("log.server.http.host")
	}
	if cfg.Server.Http.Port != DefaultServerHttpPort {
		t.Fatal("log.server.http.port")
	}
}

func TestConfig(t *testing.T) {
	var cfg Config
	var err error
	cwd, err := os.Getwd()
	if err == nil {
		cwd += "/"
	}
	if cfg, err = ReadConfig(cwd + "examples/config/full.toml"); err != nil {
		t.Fatal(err.Error())
	}
	if cfg.Log.Path != "gogo.log" {
		t.Fatal("log.path")
	}
	if cfg.Log.Level != "debug" {
		t.Fatal("log.level")
	}
	if cfg.Server.Username != "admin" {
		t.Fatal("server.username")
	}
	if cfg.Server.Password != "verytopsecret" {
		t.Fatal("server.password")
	}
	if cfg.Server.Http.Host != "localhost" {
		t.Fatal("server.http.host")
	}
	if cfg.Server.Http.Port != 8181 {
		t.Fatal("server.http.port")
	}
	if len(cfg.Tasks) != 2 {
		t.Fatal("tasks")
	} else {
		for k,v := range cfg.Tasks {
			if v.Cmd == "" {
				t.Fatal("tasks[" + strconv.Itoa(k) + "].cmd")
			}
		}
	}
}
