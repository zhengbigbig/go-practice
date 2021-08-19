package main

import "log"

func main() {
	//log.Println("normal log")
	//v := "normal"
	//log.Printf("%s log \n",v)
	//log.Fatalln("fatal log")
	//log.Panicln("panic log")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("log")
	log.SetPrefix("[prefix]")
	log.Println("normal log")
}
