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

var log = NewLog()

func TestNew(t *testing.T) {
	// u := NewLog()
	log.Error("test")
	log.Warn("1")
	log.Fatal("2")
}

