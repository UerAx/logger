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
	"os"
	"sync/atomic"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var (
	level        = Trace
	flag = Ldate | Ltime | Llongfile | Lmsgprefix
)


type Ulog struct {
	*Logger
	level int
}

func NewLog() *Ulog {
	return &Ulog{new(os.Stdout, "", flag),level,}
}

func NewErrLog() *Ulog {
	return &Ulog{new(os.Stderr, "", flag),level,}
}


func (t *Ulog) SetLevel(level int) {
	t.level = level
}

func (t *Ulog) SetOut(w io.Writer) {
	t.SetOutput(w)
}

func (l *Ulog) OutFile(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  	if err != nil {
    	l.Error("打开日志文件失败：", err)
	}
	l.SetOutput(io.MultiWriter(f))
}

func (l *Ulog) Trace(v... interface{}) {
	if l.level <= TRACE {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintln(v...), "[TRACE] ")
	}
}

func (l *Ulog) Tracef(format string, v ...interface{}) {
	if l.level <= INFO {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintf(format, v...), "[TRACE] ")
	}
}


func (l *Ulog) Info(v... interface{}) {
	if l.level <= INFO {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintln(v...), "[INFO] ")
	}
}

func (l *Ulog) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintf(format, v...), "[INFO] ")
	}
}

func (l *Ulog) Warn(v... interface{}) {
	if l.level <= WARN {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintln(v...), "[WARN] ")
	}
}

func (l *Ulog) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintf(format, v...), "[WARN] ")
	}
}

func (l *Ulog) Error(v... interface{}) {
	if l.level <= ERROR {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintln(v...), "[ERROR] ")
	}
}

func (l *Ulog) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		if atomic.LoadInt32(&l.isDiscard) != 0 {
			return
		}
		l.Output(2, fmt.Sprintf(format, v...), "[ERROR] ")
	}
}

func (l *Ulog) Fatal(v... interface{}) {
	if l.level <= FATAL {
		l.Output(2, fmt.Sprintln(v...), "[FATAL] ")
	}
	os.Exit(1)
}

func (l *Ulog) Fatalf(format string, v ...interface{}) {
	if l.level <= FATAL {
		l.Output(2, fmt.Sprintf(format, v...), "[FATAL] ")
	}
	os.Exit(1)
}

func (l *Ulog) Panic(v... interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(2, s, "[PANIC] ")
	panic(s)
}

func (l *Ulog) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(2, s, "[PANIC]")
	panic(s)
}
