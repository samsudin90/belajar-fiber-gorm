package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	ProjectRootPathh = filepath.Join(filepath.Dir(b), "../")
)
