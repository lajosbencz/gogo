package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-c path/to/cfg]\n", os.Args[0])
}


func main() {
	logHandlers := multi.New(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)
	log.SetHandler(logHandlers)

	binary, err := exec.LookPath("ping")
	if err != nil {
		panic(err)
	}
	fmt.Println("Resolved binary: " + binary)
	cmd := exec.Command(binary, "google.com", "-n", "3")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	/*
	var cfgPath string
	fs := flag.NewFlagSet("gogo", flag.ExitOnError)
	fs.StringVar(&cfgPath, "c", "gogo.toml", "Path config file")
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

	tasks := make([]Program, len(cfg.Programs))
	for _,v := range cfg.Programs {
		p := NewProgram(v)
		p.Start()
		tasks = append(tasks, p)
	}

	for _, v := range tasks {
		v.Cmd.Wait()
		fmt.Println(v.Cmd.CombinedOutput())
	}
	os.Exit(0)
	*/
}
