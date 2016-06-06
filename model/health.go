package model

type Health struct {
	Status  bool       `toml:"-"       json:"status"`
	Process []*Process `toml:"process" json:"process"`
}

func (h *Health) IsHealth() bool {
	h.Status = true

	for _, v := range h.Process {
		if !v.IsProcess() {
			h.Status = false
		}
	}

	return h.Status
}
