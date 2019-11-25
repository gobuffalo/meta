package meta

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/gobuffalo/here"
)

func Marshal(w io.Writer, a *App) error {
	return toml.NewEncoder(w).Encode(a)
}

func Unmarshal(info here.Info, b []byte) (*App, error) {
	a := &App{
		Info: info,
	}

	md, err := toml.Decode(string(b), a)
	if err != nil {
		return nil, err
	}
	if len(md.Undecoded()) == 0 {
		return a, nil
	}

	oa := &oldApp{}
	md, err = toml.Decode(string(b), oa)
	if err != nil {
		return nil, err
	}

	a, err = construct(info)
	if err != nil {
		return nil, err
	}

	if len(oa.Bin) != 0 {
		a.Bin = oa.Bin
	}

	if len(oa.VCS) != 0 {
		a.VCS = oa.VCS
	}

	a.With["pop"] = oa.WithPop
	a.With["sqlite"] = oa.WithSQLite
	a.With["webpack"] = oa.WithWebpack
	a.With["nodejs"] = oa.WithNodeJs
	a.With["yarn"] = oa.WithYarn

	a.As["web"] = oa.AsWeb
	a.As["api"] = oa.AsAPI

	return a, nil
}

type oldApp struct {
	Name        string `json:"name" toml:"name"`
	Bin         string `json:"bin" toml:"bin"`
	VCS         string `json:"vcs" toml:"vcs"`
	WithPop     bool   `json:"with_pop" toml:"with_pop"`
	WithSQLite  bool   `json:"with_sqlite" toml:"with_sqlite"`
	WithDep     bool   `json:"with_dep" toml:"with_dep"`
	WithWebpack bool   `json:"with_webpack" toml:"with_webpack"`
	WithNodeJs  bool   `json:"with_nodejs" toml:"with_nodejs"`
	WithYarn    bool   `json:"with_yarn" toml:"with_yarn"`
	WithDocker  bool   `json:"with_docker" toml:"with_docker"`
	WithGrifts  bool   `json:"with_grifts" toml:"with_grifts"`
	AsWeb       bool   `json:"as_web" toml:"as_web"`
	AsAPI       bool   `json:"as_api" toml:"as_api"`
}
