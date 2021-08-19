package main

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(funcName,path.Base(file))
	// path.Base拿到最后路径
}

func main() {
	f()
}
