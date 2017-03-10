package phck

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type Process struct {
	Label   string         `toml:"label"    json:"label,omitempty"`
	Running bool           `toml:"-"        json:"running"`
	PIDFile string         `toml:"pidfile"  json:"pidfile"`
	Message string         `toml:"-"        json:"message,omitempty"`
	Error   error          `toml:"-"        json:"-"`
	Detail  *ProcessDetail `toml:"-"        json:"detail,omitempty"`
}

type ProcessDetail struct {
	Name      string  `json:"name,omitempty"`
	PID       int32   `json:"pid,omitempty"`
	PPID      int32   `json:"ppid,omitempty"`
	User      string  `json:"user,omitempty"`
	Command   string  `json:"command,omitempty"`
	Stat      string  `json:"stat,omitempty"`
	Thread    int32   `json:"thread,omitempty"`
	UseMemory float32 `json:"use_memory,omitempty"`
}

func (p *Process) IsProcess() bool {
	p.Running = true
	p.Error = nil
	p.Message = ""
	p.Detail = &ProcessDetail{}

	if p.Detail.PID, p.Error = ReadPIDFile(p.PIDFile); p.Error != nil {
		p.Message = "PID file is not open"
		p.Running = false
		return p.Running
	}

	if ok, _ := process.PidExists(p.Detail.PID); !ok {
		p.Message = "PID not found"
		p.Error = errors.New(p.Message)
		p.Running = false
		return p.Running
	}

	d, _ := process.NewProcess(p.Detail.PID)

	p.Detail.Name, p.Error = d.Name()
	p.Detail.User, p.Error = d.Username()
	p.Detail.PPID, p.Error = d.Ppid()
	p.Detail.Command, p.Error = d.Cmdline()
	p.Detail.Thread, p.Error = d.NumThreads()
	p.Detail.UseMemory, p.Error = d.MemoryPercent()
	p.Detail.Stat, p.Error = d.Status()

	if strings.Contains(p.Detail.Stat, "Z") {
		p.Message = "Process is zombie"
		p.Error = errors.New(p.Message)
		p.Running = false
		return p.Running
	}

	if strings.Contains(p.Detail.Stat, "X") {
		p.Message = "Process is dead"
		p.Error = errors.New(p.Message)
		p.Running = false
		return p.Running
	}

	return p.Running
}

func ReadPIDFile(filename string) (int32, error) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(b)))

	if err != nil {
		return 0, err
	}

	return int32(pid), nil
}
