/*
 * @Author: ww
 * @Date: 2022-06-29 06:36:32
 * @Description:
 * @FilePath: /ulog/ulog.go
 */

package ulog

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Level uint

type Ulog struct {
	info *log.Logger
	warn *log.Logger
	err *log.Logger
	Level Level
}

func New() *Ulog {
	return &Ulog{
		log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		INFO,
	}
}

func (l *Ulog) OutFile(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  	if err != nil {
    	l.Error("打开日志文件失败：", err)
	}
	l.info.SetOutput(io.MultiWriter(f, os.Stdout))
	l.warn.SetOutput(io.MultiWriter(f, os.Stdout))
	l.err.SetOutput(io.MultiWriter(f, os.Stderr))
}

func (l *Ulog) Info(v... interface{}) {
	if l.Level <= INFO {
		l.info.Println(v...)
	}
}

func (l *Ulog) Infof(format string, v ...interface{}) {
	if l.Level <= INFO {
		l.info.Println(fmt.Sprintf(format, v...))
	}
}

func (l *Ulog) Warn(v... interface{}) {
	if l.Level <= WARN {
		l.warn.Println(v...)
	}
}

func (l *Ulog) Warnf(format string, v ...interface{}) {
	if l.Level <= WARN {
		l.warn.Println(fmt.Sprintf(format, v...))
	}
}

func (l *Ulog) Error(v... interface{}) {
	if l.Level <= ERROR {
		l.err.Println(v...)
	}
}

func (l *Ulog) Errorf(format string, v ...interface{}) {
	if l.Level <= ERROR {
		l.err.Println(fmt.Sprintf(format, v...))
	}
}

func (l *Ulog) Fatal(v... interface{}) {
	if l.Level <= FATAL {
		l.err.Fatalln(v...)
	}
}

func (l *Ulog) Fatalf(format string, v ...interface{}) {
	if l.Level <= FATAL {
		l.err.Fatalln(fmt.Sprintf(format, v...))
	}
}

func (l *Ulog) Panic(v... interface{}) {
	l.Error(v...)
}

func (l *Ulog) Panicf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}
