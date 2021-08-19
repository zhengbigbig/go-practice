package main

import (
	"fmt"
	zbb "zhengbigbig.com/06-package/calc" // zbb 是别名
)

//import  "zhengbigbig.com/06-package/calc"

func init() {
	fmt.Println("auto")
}

func main() {
	ret := zbb.Add(10,20)
	fmt.Println(ret)
}