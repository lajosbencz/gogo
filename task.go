package main

import (
	"os/exec"
	"log"
	"time"
)

const (
	TaskStatusStopped = iota
	TaskStatusStarted
	TaskStatusRetrying
	TaskStatusRunning
	TaskStatusFailed
)

type Task struct {
	Key string
	Pid int
	Since time.Time
	Status int
	Retries uint
	Config ConfigTasks
	Cmd *exec.Cmd
}

func NewTask(key string, cfg ConfigTasks) Task {
	cmd := exec.Command(cfg.Cmd, cfg.Args...)
	p := Task{
		Key: key,
		Pid: cmd.Process.Pid,
		Status: TaskStatusStopped,
		Retries: 0,
		Since: nil,
		Config: cfg,
		Cmd: cmd,
	}
	return p
}

func (p *Task) Start() error {
	p.Since = time.Now()
	p.Status = TaskStatusStarted
	if err := p.Cmd.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Task) Stop() error {
	p.Status = TaskStatusStopped
	p.Since = time.Now()
	return nil
}

func (p *Task) Restart() error {
	p.Stop()
	p.Start()
	return nil
}

func (p *Task) running() error {
	p.Status = TaskStatusRunning
	return nil
}

func (p *Task) retry() error {
	p.Status = TaskStatusRetrying
	return nil
}

func (p *Task) failed() error {
	p.Status = TaskStatusRetrying
	p.Since = time.Now()
	return nil
}