package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/containous/traefik/log"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-c path/to/cfg]\n", os.Args[0])
}

func main() {

	binary, err := exec.LookPath("ping")
	if err != nil {
		panic(err)
	}
	log.Debug("Resolved binary: " + binary)
	cmd := exec.Command(binary, "google.com", "-n", "3")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	/*
	var cfgPath string
	fs := flag.NewFlagSet("gogo", flag.ExitOnError)
	fs.StringVar(&cfgPath, "c", "gogo", "Path config file")
	fs.Usage = usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfgBase, cfgFile := path.Split(cfgPath)

	fmt.Println("cfgBase: ",cfgBase, "cfgFile: ",cfgFile)

	cfg, err := ReadConfig(cfgFile, []string{cfgBase})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	programs := make([]Program, len(cfg.Programs))
	for _,v := range cfg.Programs {
		p := NewProgram(v)
		p.Start()
		programs = append(programs, p)
	}

	for _, v := range programs {
		v.Cmd.Wait()
		fmt.Println(v.Cmd.CombinedOutput())
	}
	os.Exit(0)
	*/
}
