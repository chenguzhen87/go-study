package main

import (
	"log"
	"log/syslog"
)

func main() {
	testSysLog()
}

func testSysLog() {
	err := syslog.Dial("", "",syslog.LOG_ERR, "Saturday")
	if err != nil {
		log.Fatal(err)
	}

	sysLog.Emerg("Hello world!")
}
