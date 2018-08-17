package main

import (
	"runtime"

	"github.com/funayman/aomori-library/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.Execute()
}
