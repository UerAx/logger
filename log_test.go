package logger

import (
	"fmt"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	os.Chdir("..")
	InitLogger(1)
	x := "1"
	Log("debug", "debug")
	Logf("debug %s %s", x, x)
	Debug("debug", "debug")
	Debugf("debug %s %s", x, x)
	Infof("info %s %s", x, x)
	Info("info","info")
	Warn("warm", "warm", "warm")
	Warnf("info %s %d", x, 1)
	Error("error", "error", "error")
	Errorf("error %s %d", x, 1)
	// Fatal("fa", "warm", "warm")
	fmt.Println(1111)
}
