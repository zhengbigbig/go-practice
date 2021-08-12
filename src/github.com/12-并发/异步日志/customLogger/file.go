package customLogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志相关

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

func NewFileLog(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level: level, filePath: fp, fileName: fn, maxFileSize: maxSize,
		logChan: make(chan *logMsg, 50000),
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file open failed,err:", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file open failed,err:", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启goroutine去往文件写日志
	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) isOverMaxSize(fileObj *os.File) bool {
	// 1. 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	fmt.Println(fileInfo, "fileInfo", err)
	if err != nil {
		fmt.Println("get file info failed,err:", err)
		return false
	}
	size := fileInfo.Size() // B
	return size >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要日志切割
	bkStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed,err:", err)
		return nil, err
	}
	oldNamePath := path.Join(f.filePath, fileInfo.Name())     // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", oldNamePath, bkStr) // 拼接一个备份名
	// 1. 关闭当前日志文件
	file.Close()
	// 2. 备份一下 rename
	os.Rename(oldNamePath, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(oldNamePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file open failed,err:", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象返回
	return fileObj, nil

}

func (f *FileLogger) checkSize(lv interface{}) {
	tg := f.fileObj
	if lv != nil {
		tg = f.errFileObj
	}
	if f.isOverMaxSize(tg) {
		newFile, err := f.splitFile(tg)
		if err != nil {
			fmt.Println("split file err:", err)
			return
		}
		if lv != nil {
			f.errFileObj = newFile
		} else {
			f.fileObj = newFile
		}
	}
}

func (f *FileLogger) writeLogBackground() {
	for {
		f.checkSize(nil)
		// 避免取不到阻塞
		select {
		case logTemp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTemp.timestamp, getLogString(logTemp.Level),
				logTemp.fileName, logTemp.funcName, logTemp.line, logTemp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTemp.Level >= ERROR {
				f.checkSize("err")
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) // 从里到外
		// 先把日志发送到通道中
		logTemp := &logMsg{
			Level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select { // 避免满
		case f.logChan <- logTemp:
		default:
			// 把日志就丢掉，保证不出现阻塞
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
