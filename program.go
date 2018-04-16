package main

import (
	"os/exec"
	"log"
)

type Program struct {
	Config ConfigProgram
	Cmd *exec.Cmd
}

func NewProgram(cfg ConfigProgram) Program {
	cmd := exec.Command(cfg.Cmd, cfg.Args...)
	p := Program{Config: cfg, Cmd: cmd}
	return p
}

func (p *Program) Start() error {
	if err := p.Cmd.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Program) Stop() error {
	return nil
}

func (p *Program) Restart() error {
	return nil
}