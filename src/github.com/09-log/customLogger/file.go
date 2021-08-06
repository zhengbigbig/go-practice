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
}

func NewFileLog(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level: level, filePath: fp, fileName: fn, maxFileSize: maxSize,
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
	if err != nil {
		fmt.Println("get file info failed,err:", err)
		return false
	}
	size := fileInfo.Size() // B
	return size >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要日志切割
	// 1. 关闭当前日志文件
	file.Close()
	// 2. 备份一下 rename
	bkStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed,err:", err)
		return nil, err
	}
	oldNamePath := path.Join(f.filePath, fileInfo.Name())     // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", oldNamePath, bkStr) // 拼接一个备份名
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

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) // 从里到外
		f.checkSize(nil)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			f.checkSize("err")
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
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
