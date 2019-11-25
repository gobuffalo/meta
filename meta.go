package meta

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gobuffalo/here"
)

func NewDir(dir string) (*App, error) {
	info, err := here.Dir(dir)
	if err != nil {
		return nil, err
	}
	return New(info)
}

func New(info here.Info) (*App, error) {
	a := &App{
		Info: info,
		With: With{},
		As:   As{},
	}
	a.Bin = filepath.Join("bin", a.Info.Name)
	if runtime.GOOS == "windows" {
		a.Bin += ".exe"
	}

	for k, v := range vcses {
		_, err := os.Stat(filepath.Join(info.Module.Dir, v))
		if err != nil {
			continue
		}
		a.VCS = k
		break
	}

	return a, nil
}

var vcses = map[string]string{
	"git": ".git",
	"bzr": ".bzr",
	"svn": ".svn",
}
