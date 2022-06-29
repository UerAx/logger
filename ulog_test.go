/*
 * @Author: ww
 * @Date: 2022-06-29 06:36:32
 * @Description:
 * @FilePath: /ulog/ulog_test.go
 */

package ulog

import (
	"testing"
)

func TestNew(t *testing.T) {
	log := New()
	log.OutFile("./log.log")
	log.Level = 3
	log.Info("test", 1)
	log.Warn("test",2)
	log.Error("test", 3)
	log.Fatal("test4")
	log.Panic("test56")
}

