package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/andygeiss/meridian59-build/patcher/internal/patcher"
)

func main() {
	// retrieve all the patch files recursively
	os.Chdir("client")
	var patchFiles []patcher.PatchFile
	files := patcher.GetFiles(".")
	for _, path := range files {
		patchFiles = append(patchFiles, patcher.NewPatchFile(path))
	}
	// encode it as a json file
	target, _ := os.Create(filepath.Join("patchinfo.txt"))
	_ = json.NewEncoder(target).Encode(patchFiles)
	target.Close()
}
