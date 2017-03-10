package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	defaultCfgName = "phck.conf"
	version        = "1.0"
)

var (
	cfgName   = flag.String("c", defaultCfgName, "set cfgiguration file")
	isHelp    = flag.Bool("h", false, "this help")
	isVersion = flag.Bool("v", false, "show this build version")
	isCLI     = flag.Bool("cli", false, "CLI mode")
)

func main() {
	os.Exit(exitcode(run()))
}

func exitcode(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return 1
	}

	return 0
}

func run() error {
	flag.Parse()

	if *isVersion {
		fmt.Println("v" + version)
		return nil
	}

	if *isHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		return nil
	}

	cfg := newConfig()

	if err := cfg.read(*cfgName); err != nil {
		return err
	}

	if *isCLI {
		cfg.Health.IsHealth()

		b, err := json.Marshal(cfg.Health)

		if err != nil {
			return err
		}

		fmt.Printf("%s", b)

		return nil
	}

	srv := newServer(cfg)

	if len(cfg.PIDFile) > 0 {
		if err := srv.writePIDFIle(cfg.PIDFile); err != nil {
			return err
		}
	}

	if err := srv.listenAndServe(cfg.Listen); err != nil {
		return err
	}

	return nil
}
