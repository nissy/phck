package model

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type Process struct {
	Label   string        `toml:"Label"   json:"label"`
	Running bool          `toml:"-"       json:"running"`
	PIDFile string        `toml:"PIDFile" json:"pidfile"`
	Message string        `toml:"-"       json:"message,omitempty"`
	Error   error         `toml:"-"       json:"-"`
	Detail  ProcessDetail `toml:"-"       json:"detail,omitempty"`
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
	p.Running = false
	p.Error = nil
	p.Message = ""
	p.Detail = ProcessDetail{}

	if p.Detail.PID, p.Error = ReadPIDFile(p.PIDFile); p.Error != nil {
		p.Message = "PID file is not open"
		return p.Running
	}

	if ok, _ := process.PidExists(p.Detail.PID); !ok {
		p.Message = "PID not found"
		return p.Running
	}

	p.SetDetail(p.Detail.PID)
	return p.Running
}

func (p *Process) SetDetail(pid int32) {
	p.Running = true
	d, _ := process.NewProcess(pid)
	p.Detail.Name, _ = d.Name()
	p.Detail.User, _ = d.Username()
	p.Detail.PPID, _ = d.Ppid()
	p.Detail.Command, _ = d.Cmdline()
	p.Detail.Thread, _ = d.NumThreads()
	p.Detail.UseMemory, _ = d.MemoryPercent()
	p.Detail.Stat, _ = d.Status()

	if strings.Contains(p.Detail.Stat, "Z") {
		p.Error = errors.New("Process is zombie")
		p.Running = false
	}

	if strings.Contains(p.Detail.Stat, "X") {
		p.Error = errors.New("Process is dead")
		p.Running = false
	}
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
