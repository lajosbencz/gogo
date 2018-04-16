package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"os"
	"github.com/pkg/errors"
)

type Manager struct {
	Running bool
	Config Config
	tasks []Task
	logWriter *os.File
}

func New(config *Config) (*Manager, error) {
	m := Manager{}
	if err := m.UpdateConfig(config); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Manager) ResizeTask(key string, newSize uint) {

}

func (m *Manager) SetLogPath(path string) error {
	if m.logWriter != nil {
		m.logWriter.Close()
	}
	var err error
	if m.logWriter, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	handlers := multi.New(text.New(os.Stdout), text.New(m.logWriter))
	log.SetHandler(handlers)
	return nil
}

func (m *Manager) SetLogLevel(level string) error {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "", "error":
		log.SetLevel(log.DebugLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		return errors.New("Invalid log level: " + level)
	}
	return nil
}

func (m *Manager) UpdateConfig(config *Config) error {
	if m.Config.Log.Path != config.Log.Path {
		if err := m.SetLogPath(config.Log.Path); err !=nil {
			return err
		}
	}
	if m.Config.Log.Level != config.Log.Level {
		if err := m.SetLogLevel(config.Log.Level); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) UpdateTaskConfig(key string, config *ConfigTaskItem) error {
	cfgTask, ok := m.Config.Tasks[key]
	if !ok {
		m.Config.Tasks[key] = ConfigTaskItem{}
	} else {
		m.Config.Tasks[key] = cfgTask
	}
	if m.Config.Tasks[key].Cmd != config.Cmd {
		//m.SetTaskCmd(config.cmd)
	}
	return nil
}

func (m *Manager) Run() error {
	m.SetLogPath(m.Config.Log.Path)
	m.SetLogLevel(m.Config.Log.Level)
	// @todo
	m.Running = true
	return nil
}