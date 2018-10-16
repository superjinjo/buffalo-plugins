package plugdeps

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var heroku = Plugin{
	Binary: "buffalo-heroku",
	GoGet:  "github.com/gobuffalo/buffalo-heroku",
}

func Test_ConfigPath(t *testing.T) {
	r := require.New(t)

	x := ConfigPath(meta.App{Root: "foo"})
	r.Equal(x, filepath.Join("foo", "config", "buffalo-plugins.toml"))
}

func Test_List_Off(t *testing.T) {
	r := require.New(t)

	app := meta.App{}
	plugs, err := List(app)
	r.Error(err)
	r.Equal(errors.Cause(err), ErrMissingConfig)
	r.Len(plugs.List(), 0)
}

func Test_List_On(t *testing.T) {
	r := require.New(t)

	app := meta.New(os.TempDir())

	p := ConfigPath(app)
	r.NoError(os.MkdirAll(filepath.Dir(p), 0755))
	f, err := os.Create(p)
	r.NoError(err)
	f.WriteString(eToml)
	r.NoError(f.Close())

	plugs, err := List(app)
	r.NoError(err)
	r.Len(plugs.List(), 2)
}

const eToml = `[[plugin]]
  binary = "buffalo-heroku"
  go_get = "github.com/gobuffalo/buffalo-heroku"

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"`