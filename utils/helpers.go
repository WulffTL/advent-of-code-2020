package utils

import (
	"path"
	"runtime"
)

// GetRootPath returns the root directory of the project
func GetRootPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot get file into for env")
	}

	return path.Dir(path.Dir(filename))
}
