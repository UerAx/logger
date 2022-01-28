package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)


// 按照 年-月-日_时:分:秒 格式形式返回时间字符串
func GetCurrentTimeString() string {
	return time.Now().Format("2006-01-02")
}

func outFile(file string) (io.Writer, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s-%s", file, GetCurrentTimeString()), os.O_APPEND|os.O_CREATE|os.O_RDWR, 644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func Log(v... interface{}) {
	l.Log(fmt.Sprintf("[LOG] %s", fmt.Sprintln(v...)))
}

func Logf(format string, v ...interface{}) {
	l.Log(fmt.Sprintf("[LOG] %s", fmt.Sprintf(format, v...)))
}

func Debug(v... interface{}) {
	if l.Level <= DEBUG {
		l.Log(fmt.Sprintf("[DEBUG] %s", fmt.Sprintln(v...)))
	}
}

func Debugf(format string, v ...interface{}) {
	if l.Level <= DEBUG {
		l.Log(fmt.Sprintf("[DEBUG] %s", fmt.Sprintf(format, v...)))
	}
}

func Info(v... interface{}) {
	if l.Level <= INFO {
		l.Log(fmt.Sprintf("[INFO] %s", fmt.Sprintln(v...)))
	}
}

func Infof(format string, v ...interface{}) {
	if l.Level <= INFO {
		l.Log(fmt.Sprintf("[INFO] %s", fmt.Sprintf(format, v...)))
	}
}

func Warn(v... interface{}) {
	if l.Level <= WARN {
		l.Log(fmt.Sprintf("[WARM] %s", fmt.Sprintln(v...)))
	}
}

func Warnf(format string, v ...interface{}) {
	if l.Level <= WARN {
		l.Log(fmt.Sprintf("[WARM] %s", fmt.Sprintf(format, v...)))
	}
}

func Error(v... interface{}) {
	if l.Level <= ERROR {
		l.Log(fmt.Sprintf("[ERROR] %s", fmt.Sprintln(v...)))
	}
}

func Errorf(format string, v ...interface{}) {
	if l.Level <= ERROR {
		l.Log(fmt.Sprintf("[ERROR] %s", fmt.Sprintf(format, v...)))
	}
}

func Fatal(v... interface{}) {
	if l.Level <= FATAL {
		l.Fatal(fmt.Sprintf("[FATAL] %s", fmt.Sprintln(v...)))
	}
}

func Fatalf(format string, v ...interface{}) {
	if l.Level <= FATAL {
		l.Fatal(fmt.Sprintf("[FATAL] %s", fmt.Sprintf(format, v...)))
	}
}