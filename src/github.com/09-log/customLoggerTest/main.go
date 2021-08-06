package main

import logger "github.com/09-log/customLogger"

// 测试自定义日志库
func main() {
	//log := customLogger.NewLog("debug")
	log := logger.NewFileLog("debug", "/Users/zhengzhiheng/Desktop/go-practice/", "log", 1024*1024)
	log.Debug("debugger日志 %s\n", "sss")
	log.Info("info日志")
	log.Warning("warning日志")
	log.Error("error日志")
	log.Fatal("fatal日志")
}
