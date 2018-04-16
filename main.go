package main

import (
	"fmt"
	"os"
	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"flag"
	"crypto/sha256"
	"encoding/hex"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-c path/to/cfg]\n", os.Args[0])
}


func main() {

	s := []byte("secret")
	h := md5.New()
	h.Write(s)
	fmt.Println("md5: " + hex.EncodeToString(h.Sum(nil)))
	h = sha1.New()
	h.Write(s)
	fmt.Println("sha1: " + hex.EncodeToString(h.Sum(nil)))
	h = sha256.New()
	h.Write(s)
	fmt.Println("sha256: " + hex.EncodeToString(h.Sum(nil)))
	h = sha512.New()
	h.Write(s)
	fmt.Println("sha512: " + hex.EncodeToString(h.Sum(nil)))
	return

	var cfgPath string
	fs := flag.NewFlagSet("gogo", flag.ExitOnError)
	fs.StringVar(&cfgPath, "c", "gogo.toml", "Path config file")
	fs.Usage = usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		usage()
		os.Exit(1)
	}

	cfg, err := ReadConfig(cfgPath)

	if err != nil {
		panic(err)
	}

	if cfg.Log.Path != "" {
		fw, err := os.Open(cfg.Log.Path)
		defer fw.Close()
		if err != nil {
			panic(err)
		}
		th := text.New(fw)
		logHandlers := multi.New(text.New(os.Stdout), th)
		log.SetHandler(logHandlers)
	} else {
		logHandlers := multi.New(text.New(os.Stdout))
		log.SetHandler(logHandlers)
	}

	switch cfg.Log.Level {
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
	}

	manager := Manager{Config: cfg}
	manager.Run()


	/*

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
