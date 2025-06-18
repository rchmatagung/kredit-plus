package general

import (
	"path/filepath"
	"runtime"
)

var (
	//* Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	//* Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)

const (
	FullTimeFormat        string = "2006-01-02 15:04:05"
	DisplayDateTimeFormat string = "02 Jan 2006 15:04:05"
	DateFormat            string = "2006-01-02"
)
