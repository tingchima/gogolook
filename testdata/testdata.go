// Package testdata provides
package testdata

import (
	"path/filepath"
	"runtime"
)

var base string

const (
	TestDataTasks = "tasks.yml"
)

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	base = filepath.Dir(currentFile)
}

func Path(file string) string {
	return filepath.Join(base, file)
}
