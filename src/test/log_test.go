package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	log.Print("testPrint")
}

func TestPrintf(t *testing.T) {
	log.Printf("testPrintf")
}

func TestFatal(t *testing.T) {
	log.Fatal("Fatal")
}

func TestLogFile(t *testing.T) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}