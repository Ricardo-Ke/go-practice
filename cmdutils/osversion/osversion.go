package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		panic("not support windows plaform")
	} else {
		fmt.Println(runtime.GOOS)
	}
}
