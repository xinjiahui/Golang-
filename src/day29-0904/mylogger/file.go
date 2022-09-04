package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string //文件保存的路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}

	return fl
}
func (f *FileLogger) initFile() error {
	fullfileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullfileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullfileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error log file failed,err:%v", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
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
func (f *FileLogger) Trace(format string, a ...interface{}) {

	f.log(TRACE, format, a...)

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

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
