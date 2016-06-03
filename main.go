package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	"github.com/ngc224/hck/config"
	"github.com/ngc224/hck/server"
)

const (
	ApplicationName = "HCK"
	ConfigFilePath  = "./hck.conf"
)

type Command struct {
	Options *Options
	Args    []string
}

type Options struct {
	Help    bool `short:"h" long:"help"    description:"Show this help message"`
	Version bool `short:"v" long:"version" description:"Show this build version"`
}

func NewCommand() (*Command, error) {
	opts := &Options{}
	p := flags.NewParser(opts, flags.None)
	p.Name = ApplicationName
	p.Usage = "[OPTIONS] CONFIG_FILE"
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

	return ConfigFilePath
}

func main() {
	os.Exit(_main())
}

func _main() int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd, err := NewCommand()

	if err != nil {
		return PrintError(err)
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
		return PrintError(err)
	}

	server.NewServer(c).Run()

	return 0
}

func PrintError(err error) int {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	return 1
}
