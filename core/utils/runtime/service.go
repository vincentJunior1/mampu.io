package runtime

import (
	"fmt"
	"runtime"
)

func GetCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}
