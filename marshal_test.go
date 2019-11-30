package meta

import (
	"bytes"
	"os"
	"testing"

	"github.com/gobuffalo/here"
	"github.com/stretchr/testify/require"
)

func Test_App_Toml(t *testing.T) {
	r := require.New(t)

	info, err := here.Dir(".")
	r.NoError(err)

	a := App{
		Info: info,
		With: With{
			"pop": true,
			"dep": false,
		},
		As: As{
			"web": true,
			"api": false,
		},
	}

	bb := &bytes.Buffer{}

	r.NoError(Marshal(bb, &a))

	a2, err := Unmarshal(info, bb.Bytes())
	r.NoError(err)

	r.Equal(a, *a2)
}

func Test_App_Old_Toml(t *testing.T) {
	r := require.New(t)

	old := `name = "coke"
bin = "bin/coke"
vcs = "git"
with_pop = true
with_sqlite = false
with_dep = false
with_webpack = true
with_nodejs = true
with_yarn = true
with_docker = true
with_grifts = true
as_web = true
as_api = false`

	info, err := here.Dir(".")
	r.NoError(err)

	a, err := Unmarshal(info, []byte(old))
	r.NoError(err)

	Marshal(os.Stdout, a)

	r.Equal(info, a.Info)
	r.Equal("bin/coke", a.Bin)
	r.Equal("git", a.VCS)

	r.True(a.With["pop"])
	r.True(a.With["webpack"])
	r.True(a.With["nodejs"])
	r.True(a.With["yarn"])
	r.False(a.With["sqlite"])

	r.True(a.As["web"])
	r.False(a.As["api"])
}
