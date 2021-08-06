package main

import "github.com/09-log/customLogger"

// 测试自定义日志库
func main() {
	log := customLogger.NewLog("debug")
	log.Debug("debugger日志")
	log.Info("info日志")
	log.Warning("warning日志")
	log.Error("error日志")
	log.Fatal("fatal日志")
}
