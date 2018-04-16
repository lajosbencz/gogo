package main

import (
	"testing"
	"strconv"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("full", []string{"./examples/config"})
	if err != nil {
		panic("Config test failed: " + err.Error())
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
	if cfg.Server.Http.Enable != true {
		t.Fatal("server.http.enable")
	}
	if cfg.Server.Http.Host != "localhost" {
		t.Fatal("server.http.host")
	}
	if cfg.Server.Http.Port != 8181 {
		t.Fatal("server.http.port")
	}
	if len(cfg.Programs) != 2 {
		t.Fatal("programs")
	} else {
		for k,v := range cfg.Programs {
			if v.Name == "" {
				t.Fatal("programs[" + strconv.Itoa(k) + "].name")
			}
		}
	}
}
