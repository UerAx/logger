# ulog
基于golang自带的log组件进行改造

# usage

// init

log := NewLog()

// default level

log.SetLevel(1)

// default null

log.OutFile(FILE_PATH)

log.Info("INFO")

log.Warn("WARN")

log.Error("ERROR)

log.Fatal("FATAL")

log.Panic("PANIC")

log.Infof("%s","INFO")
...
