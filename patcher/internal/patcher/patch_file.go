package patcher

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// PatchFile ...
type PatchFile struct {
	Basepath string
	Download bool
	Filename string
	Version  int
	Length   int
	MyHash   string
}

// NewPatchFile ...
func NewPatchFile(path string) PatchFile {

	baseDir := filepath.Dir(path)
	if baseDir == "." {
		baseDir = "\\"
	} else {
		baseDir = "\\" + baseDir + "\\"
	}

	return PatchFile{
		Basepath: baseDir,
		Download: isFileDownloadable(path),
		Filename: filepath.Base(path),
		Length:   getFileLength(path),
		MyHash:   getFileHash(path),
		Version:  3,
	}
}

func isFileDownloadable(path string) bool {
	if strings.Contains(path, "configuration.xml") ||
		strings.Contains(path, "patchinfo.txt") ||
		strings.Contains(path, "metagen") ||
		strings.Contains(path, ".log") ||
		strings.Contains(path, ".php") {
		return false
	}
	return true
}

func getFileLength(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	info, _ := os.Stat(path)
	return int(info.Size())
}

func getFileHash(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return strings.ToUpper(fmt.Sprintf("%x", string(hash.Sum(nil))))
}
