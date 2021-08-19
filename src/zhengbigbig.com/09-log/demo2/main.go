package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile,err:= os.OpenFile("./xxx.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		fmt.Println("open log file failed,err:",err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
	log.Println("normal log")
	log.SetPrefix("[prefix]")
	log.Println("normal log")
}
