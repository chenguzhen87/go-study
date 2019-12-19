package test

import (
	"log/syslog"
	"testing"
)

func TestSysLog(t *testing.T) {
	sysLog, _ := syslog.New(syslog.LOG_NOTICE|syslog.LOG_LOCAL0, "mmmmd")

	sysLog.Warning("Warning:helloworld")

	sysLog.Close()
}