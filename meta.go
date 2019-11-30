package meta

import (
	"bytes"
	"io/ioutil"
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
	cfg := filepath.Join(info.Module.Dir, "config", "buffalo-app.toml")
	if _, err := os.Stat(cfg); err == nil {
		b, err := ioutil.ReadFile(cfg)
		if err != nil {
			return nil, err
		}
		return Unmarshal(info, b)
	}
	return construct(info)
}

func construct(info here.Info) (*App, error) {
	a := &App{
		Info: info,
		With: With{},
		As: As{
			"web": true,
		},
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

	root := info.Module.Dir
	if _, err := os.Stat(filepath.Join(root, "package.json")); err == nil {
		a.With["nodejs"] = true
	}
	if _, err := os.Stat(filepath.Join(root, "webpack.config.js")); err == nil {
		a.With["webpack"] = true
	}
	if _, err := os.Stat(filepath.Join(root, "yarn.lock")); err == nil {
		a.With["yarn"] = true
	}
	db := filepath.Join(root, "database.yml")
	if _, err := os.Stat(db); err == nil {
		a.With["pop"] = true
		if b, err := ioutil.ReadFile(db); err == nil {
			a.With["sqlite"] = bytes.Contains(bytes.ToLower(b), []byte("sqlite"))
		}
	}
	return a, nil
}

var vcses = map[string]string{
	"git": ".git",
	"bzr": ".bzr",
	"svn": ".svn",
}
