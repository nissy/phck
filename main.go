package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	"github.com/ngc224/phck/config"
	"github.com/ngc224/phck/server"
)

const (
	ApplicationName       = "phck"
	defaultConfigFilePath = "phck.conf"
)

type Command struct {
	Options *Options
	Args    []string
}

type Options struct {
	Cli     bool   `short:"c" long:"cli"description:"CLI mode"`
	PIDFile string `long:"pidfile" description:"Set PIDFILE"`
	Help    bool   `short:"h" long:"help" description:"Show this help message"`
	Version bool   `short:"v" long:"version" description:"Show this build version"`
}

func NewCommand() (*Command, error) {
	opts := &Options{}
	p := flags.NewParser(opts, flags.None)
	p.Name = ApplicationName
	p.Usage = "[options] CONFIGFILE"
	args, err := p.Parse()

	if err != nil {
		return nil, err
	}

	if opts.Help {
		p.WriteHelp(os.Stdout)
	}

	return &Command{
		Options: opts,
		Args:    args,
	}, nil
}

func (cmd *Command) GetConfigFilePath() string {
	if len(cmd.Args) > 0 {
		return cmd.Args[0]
	}

	return defaultConfigFilePath
}

func main() {
	os.Exit(cli())
}

func cli() int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd, err := NewCommand()

	if err != nil {
		return printError(err)
	}

	if cmd.Options.Help {
		return 0
	}

	if cmd.Options.Version {
		fmt.Println(Version)
		return 0
	}

	c, err := config.NewConfig(cmd.GetConfigFilePath())

	if err != nil {
		return printError(err)
	}

	if len(cmd.Options.PIDFile) > 0 {
		c.PIDFile = cmd.Options.PIDFile
	}

	if cmd.Options.Cli {
		if !c.Health.IsHealth() {

		}

		b, err := json.Marshal(c.Health)

		if err != nil {
			return printError(err)
		}

		fmt.Printf("%s", b)
		return 0
	}

	return listenServer(c)
}

func listenServer(c *config.Config) int {
	if err := server.NewServer(c).Run(); err != nil {
		return printError(err)
	}

	return 0
}

func printError(err error) int {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	return 1
}
