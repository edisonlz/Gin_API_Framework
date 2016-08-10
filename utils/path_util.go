package utils


import (
    "path"
    "runtime"
)

func CallerSourcePath() string {
    _, callerPath, _, _ := runtime.Caller(1)
    return path.Dir(callerPath)
}