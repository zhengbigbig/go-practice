package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go worker()

	select {}
}

// simple worker
func worker() {

	strSlice := []string{}
	for {
		str := "hello world "
		strSlice = append(strSlice, str)

		time.Sleep(time.Second)
	}

}
