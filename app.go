package meta

import (
	"github.com/gobuffalo/here"
)

type App struct {
	Info here.Info `json:"-" toml:"-"`
	Bin  string    `json:"bin" toml:"bin"`
	VCS  string    `json:"vcs" toml:"vcs"`
	With With      `json:"with" toml:"with"`
	As   As        `json:"as" toml:"as"`
}

type With map[string]bool
type As map[string]bool
