package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/here"
	"github.com/gobuffalo/meta/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func fromCfg(cfg string, info here.Info) (*meta.App, error) {
	b, err := ioutil.ReadFile(cfg)
	if err != nil {
		return nil, err
	}

	return meta.Unmarshal(info, b)
}

func run() error {
	cfg := filepath.Join("config", "buffalo-app.toml")

	info, err := here.Dir(".")
	if err != nil {
		return err
	}

	var a *meta.App

	_, err = os.Stat(cfg)
	if err == nil {
		a, err = fromCfg(cfg, info)
	} else {
		a, err = meta.New(info)
	}
	if err != nil {
		return err
	}

	var w io.Writer = os.Stdout
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "-w" {
			f, err := os.Create(cfg)
			if err != nil {
				return err
			}
			w = f
			defer f.Close()
		}
	}

	meta.Marshal(w, a)
	return nil
}
