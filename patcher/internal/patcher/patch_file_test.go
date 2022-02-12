package patcher_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/meridian59-build/patcher/internal/patcher"
)

func Test_NewPatchFile(t *testing.T) {
	testCases := []struct {
		base        string
		desc        string
		patchFiles  []patcher.PatchFile
		sourceFiles []string
	}{
		{
			base: filepath.Join("testdata", "1"),
			desc: "test patch file",
			patchFiles: []patcher.PatchFile{
				{
					Basepath: "\\",
					Download: true,
					Version:  3,
					Filename: "test.bgf",
					Length:   3,
					MyHash:   "ACBD18DB4CC2F85CEDEF654FCCC4A4D8",
				},
			},
			sourceFiles: []string{"test.bgf"},
		},
	}
	cwd, _ := os.Getwd()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Chdir(tC.base)
			for i, file := range tC.sourceFiles {
				patchFile := patcher.NewPatchFile(file)
				if fmt.Sprintf("%v", patchFile) != fmt.Sprintf("%v", tC.patchFiles[i]) {
					t.Fatalf("got %v but expected %v", patchFile, tC.patchFiles[i])
				}
			}
			os.Chdir(cwd)
		})
	}
}
