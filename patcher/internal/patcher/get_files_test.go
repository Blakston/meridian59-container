package patcher_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/meridian59-build/patcher/internal/patcher"
)

func Test_GetFiles(t *testing.T) {
	testCases := []struct {
		base  string
		desc  string
		files []string
	}{
		{
			base: filepath.Join("testdata", "1"),
			desc: "test one file",
			files: []string{
				"test.bgf",
			},
		},
		{
			base: filepath.Join("testdata", "2"),
			desc: "test two files with sub dir",
			files: []string{
				filepath.Join("sub", "test2.bgf"),
				"test.bgf",
			},
		},
	}
	cwd, _ := os.Getwd()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Chdir(tC.base)
			files := patcher.GetFiles(".")
			if fmt.Sprintf("%v", tC.files) != fmt.Sprintf("%v", files) {
				t.Fatalf("got %v but expected %v", files, tC.files)
			}
			os.Chdir(cwd)
		})
	}
}
