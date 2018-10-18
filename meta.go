package meta

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gobuffalo/envy"
)

// New App based on the details found at the provided root path
func New(root string) App {
	pwd, _ := os.Getwd()
	if root == "." {
		root = pwd
	}

	// Handle symlinks
	var oldPwd = pwd
	pwd = ResolveSymlinks(pwd)
	os.Chdir(pwd)
	if runtime.GOOS != "windows" {
		// On Non-Windows OS, os.Getwd() uses PWD env var as a preferred
		// way to get the working dir.
		os.Setenv("PWD", pwd)
	}
	defer func() {
		// Restore PWD
		os.Chdir(oldPwd)
		if runtime.GOOS != "windows" {
			os.Setenv("PWD", oldPwd)
		}
	}()

	// Gather meta data
	name := filepath.Base(root)
	pp := resolvePackageName(name, pwd, modsOn)

	app := App{
		Pwd:         pwd,
		Root:        root,
		GoPath:      envy.GoPath(),
		Name:        name,
		PackagePkg:  pp,
		ActionsPkg:  pp + "/actions",
		ModelsPkg:   pp + "/models",
		GriftsPkg:   pp + "/grifts",
		WithModules: modsOn,
	}

	app.Bin = filepath.Join("bin", filepath.Base(root))

	if runtime.GOOS == "windows" {
		app.Bin += ".exe"
	}
	db := filepath.Join(root, "database.yml")
	if _, err := os.Stat(db); err == nil {
		app.WithPop = true
		if b, err := ioutil.ReadFile(db); err == nil {
			app.WithSQLite = bytes.Contains(bytes.ToLower(b), []byte("sqlite"))
		}
	}
	if _, err := os.Stat(filepath.Join(root, "Gopkg.toml")); err == nil {
		app.WithDep = true
	}
	if _, err := os.Stat(filepath.Join(root, "webpack.config.js")); err == nil {
		app.WithWebpack = true
	}
	if _, err := os.Stat(filepath.Join(root, "yarn.lock")); err == nil {
		app.WithYarn = true
	}
	if _, err := os.Stat(filepath.Join(root, "Dockerfile")); err == nil {
		app.WithDocker = true
	}
	if _, err := os.Stat(filepath.Join(root, "grifts")); err == nil {
		app.WithGrifts = true
	}
	if _, err := os.Stat(filepath.Join(root, ".git")); err == nil {
		app.VCS = "git"
	} else if _, err := os.Stat(filepath.Join(root, ".bzr")); err == nil {
		app.VCS = "bzr"
	}

	return app
}
