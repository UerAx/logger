package logger

import (
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

var (
	l *Logger

)

func InitLogger(level int) error {
	l = NewLogger(level)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return nil
}

type Logger struct {
	Level int
	OutFile string
	Prefix string
}


func NewLogger(level int) *Logger {
	return &Logger{
		Level: level,
	}
}

func (l *Logger) GetLevel() int {
	return l.Level
}

func (l *Logger) Log(v string)  {
	log.Output(2, v)
}

func (l *Logger) Fatal(v string)  {
	log.Output(2, v)
	os.Exit(1)
}



