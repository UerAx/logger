# ulog
简单的分级golang日志库

# usage

// init
log := New()

// default level
log.Level = 1

// default null
log.OutFile(FILE_PATH)

log.Info("INFO")
log.Warn("WARN")
log.Error("ERROR)
log.Fatal("FATAL")
log.Panic("PANIC")

log.Infof("%s","INFO")
...
