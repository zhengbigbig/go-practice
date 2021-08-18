package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func loginCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n",v)
		default:
			fmt.Println("sleep")
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof,"cpu",false,"turn cpu pprof on")
	flag.BoolVar(&isMemPprof,"mem",false,"turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("./cpu.pprof") // 在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n",err)
			return
		}
		pprof.StartCPUProfile(f1) // 往文件中记录CPU profile信息
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}
	for i := 0; i < 6; i++ {
		go loginCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed,err:%v\n",err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
