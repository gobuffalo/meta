package meta

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/flect/name"
)

var modsOn = (strings.TrimSpace(envy.Get("GO111MODULE", "off")) == "on")

// App represents meta data for a Buffalo application on disk
type App struct {
	Pwd         string     `json:"pwd"`
	Root        string     `json:"root"`
	GoPath      string     `json:"go_path"`
	Name        name.Ident `json:"name"`
	Bin         string     `json:"bin"`
	PackagePkg  string     `json:"package_path"`
	ActionsPkg  string     `json:"actions_path"`
	ModelsPkg   string     `json:"models_path"`
	GriftsPkg   string     `json:"grifts_path"`
	VCS         string     `json:"vcs"`
	WithPop     bool       `json:"with_pop"`
	WithSQLite  bool       `json:"with_sqlite"`
	WithDep     bool       `json:"with_dep"`
	WithWebpack bool       `json:"with_webpack"`
	WithYarn    bool       `json:"with_yarn"`
	WithDocker  bool       `json:"with_docker"`
	WithGrifts  bool       `json:"with_grifts"`
	WithModules bool       `json:"with_modules"`
}

func (a App) IsZero() bool {
	return a.String() == App{}.String()
}

func resolvePackageName(name string, pwd string, modsOn bool) string {
	result := envy.CurrentPackage()

	if filepath.Base(result) != name {
		result = path.Join(result, name)
	}

	if modsOn {
		if !strings.HasPrefix(pwd, filepath.Join(envy.GoPath(), "src")) {
			result = name
		}

		//Extract package from go.mod
		if f, err := os.Open(filepath.Join(pwd, "go.mod")); err == nil {
			if s, err := ioutil.ReadAll(f); err == nil {
				re := regexp.MustCompile("module (.*)")
				res := re.FindAllStringSubmatch(string(s), 1)

				if len(res) == 1 && len(res[0]) == 2 {
					result = res[0][1]
				}
			}
		}
	}

	return result
}

// ResolveSymlinks takes a path and gets the pointed path
// if the original one is a symlink.
func ResolveSymlinks(p string) string {
	cd, err := os.Lstat(p)
	if err != nil {
		return p
	}
	if cd.Mode()&os.ModeSymlink != 0 {
		// This is a symlink
		r, err := filepath.EvalSymlinks(p)
		if err != nil {
			return p
		}
		return r
	}
	return p
}

func (a App) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}
