package main

import logger "github.com/09-log/customLogger"

// 测试自定义日志库
var log logger.Logger // 声明一个全局的接口变量

func main() {
	log = logger.NewConsoleLog("debug")
	log = logger.NewFileLog("debug", "./", "test.log", 1024)
	for i := 0; i < 1000; i++ {
		log.Debug("debugger日志 %s\n", "sss")
		log.Info("info日志")
		log.Warning("warning日志")
		log.Error("error日志")
		log.Fatal("fatal日志")
	}
}
