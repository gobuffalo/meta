package main

import (
	"io"
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

func run() error {

	info, err := here.Dir(".")
	if err != nil {
		return err
	}

	a, err := meta.New(info)
	if err != nil {
		return err
	}

	var w io.Writer = os.Stdout
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "-w" {
			cfg := filepath.Join(info.Module.Dir, "config", "buffalo-app.toml")
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
