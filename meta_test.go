package meta

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gobuffalo/here"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	pwd, err := os.Getwd()
	r.NoError(err)

	a, err := NewDir(pwd)
	r.NoError(err)

	r.NotZero(a)

	info, err := here.Dir(pwd)
	r.NoError(err)
	r.Equal(info, a.Info)

	bin := filepath.Join("bin", "meta")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	r.Equal(bin, a.Bin)

	r.Equal("git", a.VCS)
}
