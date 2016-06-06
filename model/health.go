package model

type Health struct {
	Status     bool       `toml:"-"       json:"status"`
	StatusCode int        `toml:"-"       json:"status_code,omitempty"`
	Process    []*Process `toml:"process" json:"process"`
}

func (h *Health) IsHealth() bool {
	h.Status = true
	h.StatusCode = 0

	for _, v := range h.Process {
		if !v.IsProcess() {
			h.Status = false
		}
	}

	return h.Status
}
