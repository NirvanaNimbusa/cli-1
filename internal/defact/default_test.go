package defact

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/autarch/testify/require"
)

func Test_uniqueExes(t *testing.T) {
	tests := []struct {
		name    string
		bins    []string
		pathext string
		want    []string
	}{
		{
			"Returns same bins",
			[]string{"path1/a", "path2/b", "path3/c"},
			"",
			[]string{"path1/a", "path2/b", "path3/c"},
		},
		{
			"Returns exe prioritized",
			[]string{"path1/a.exe", "path2/a.cmd", "path3/c"},
			".exe;.cmd",
			[]string{"path1/a.exe", "path3/c"},
		},
		{
			"Returns cmd prioritized by PATH",
			[]string{"path1/a.exe", "path2/a.cmd", "path2/c"},
			".cmd;.exe",
			[]string{"path1/a.exe", "path2/c"},
		},
		{
			"Returns cmd prioritized by PATHEXT",
			[]string{"path1/a.exe", "path1/a.cmd", "path2/c"},
			".cmd;.exe",
			[]string{"path1/a.cmd", "path2/c"},
		},
		{
			"PATHEXT can be empty",
			[]string{"path1/a", "path2/b", "path3/c"},
			"",
			[]string{"path1/a", "path2/b", "path3/c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UniqueExes(tt.bins, tt.pathext)
			if err != nil {
				t.Errorf("uniqueExes error: %v", err)
				t.FailNow()
			}
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueExes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shims(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	shimFile := filepath.Join(td, "shim")
	err = createShimFile(filepath.FromSlash("/abc/def/python"), shimFile)
	require.NoError(t, err)

	require.True(t, isShimAndTargetsDir(shimFile, ""))
	require.True(t, isShimAndTargetsDir(shimFile, filepath.FromSlash("/abc/def")))
	require.False(t, isShimAndTargetsDir(shimFile, filepath.FromSlash("/some/other/dir")))
}
