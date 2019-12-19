package test

import (
	"log"
	"log/syslog"
	"testing"
)

func TestSysLog(t *testing.T) {
	sysLog, err := syslog.Dial("", "",syslog.LOG_ERR, "Saturday")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Emerg("Hello world!")
}