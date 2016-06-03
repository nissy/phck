package model

type Health struct {
	StatusCode int        `toml:"-" json:"status_code"`
	Process    []*Process `toml:"process" json:"process"`
}