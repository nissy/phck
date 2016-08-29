package phck

import (
	"os"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

type Health struct {
	DateTime string     `toml:"-"       json:"datetime"`
	HostName string     `toml:"-"       json:"hostname"`
	Status   bool       `toml:"-"       json:"status"`
	Process  []*Process `toml:"process" json:"process"`
}

func (h *Health) IsHealth() bool {
	h.Status = true
	h.DateTime = time.Now().Format(TimeFormat)

	if hostname, err := os.Hostname(); err == nil {
		h.HostName = hostname
	}

	for _, v := range h.Process {
		if !v.IsProcess() {
			h.Status = false
		}
	}

	return h.Status
}
