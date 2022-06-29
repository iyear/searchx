package global

import (
	"fmt"
	"runtime"
)

var Version string

func GetRuntime() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
